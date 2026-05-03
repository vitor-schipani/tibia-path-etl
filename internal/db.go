package internal

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClient struct {
	pool *pgxpool.Pool
}

func NewPosgresClient(ctx context.Context, connString string) (*PostgresClient, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return &PostgresClient{pool}, nil
}

func (p *PostgresClient) Close() {
	p.pool.Close()
}

func (p *PostgresClient) InsertCharacterWithDeaths(ctx context.Context, w domain.CharacterWrapper) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	c := w.Character
	_, err = tx.Exec(ctx, `
		INSERT INTO characters (name, sex, title, vocation, world, residence, account_status, unlocked_titles, level, achievement_points)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (name) DO UPDATE SET
			sex = EXCLUDED.sex,
			title = EXCLUDED.title,
			vocation = EXCLUDED.vocation,
			world = EXCLUDED.world,
			residence = EXCLUDED.residence,
			account_status = EXCLUDED.account_status,
			unlocked_titles = EXCLUDED.unlocked_titles,
			level = EXCLUDED.level,
			achievement_points = EXCLUDED.achievement_points,
			updated_at = NOW()`,
		c.Name, c.Sex, c.Title, c.Vocation, c.World, c.Residence, c.AccountStatus,
		c.UnlockedTitles, c.Level, c.AchievementPoints,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `DELETE FROM deaths WHERE character_name = $1`, c.Name)
	if err != nil {
		return err
	}

	for _, d := range w.Deaths {
		var deathID int64
		err = tx.QueryRow(ctx, `INSERT INTO deaths (character_name, level) VALUES ($1, $2) RETURNING id`, c.Name, d.Level).Scan(&deathID)
		if err != nil {
			return err
		}
		for _, k := range d.Killers {
			_, err = tx.Exec(ctx, `INSERT INTO killers (death_id, name, is_player) VALUES ($1, $2, $3)`, deathID, k.Name, k.IsPlayer)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}
