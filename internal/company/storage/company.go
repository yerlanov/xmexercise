package storage

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/yerlanov/xmexercise/internal/company"
	"github.com/yerlanov/xmexercise/pkg/postgresql"
)

func (q *Queries) Create(ctx context.Context, company company.Company) (res company.Company, err error) {
	query := `INSERT INTO public.company (name, code, country, website, phone)
   			  VALUES($1, $2, $3, $4, $5) RETURNING id, name, code, country, website, phone`

	err = q.db.QueryRow(ctx, query, company.Name, company.Code, company.Country, company.Website, company.Phone).Scan(
		&res.ID,
		&res.Name,
		&res.Code,
		&res.Country,
		&res.Website,
		&res.Phone,
	)
	if err != nil {
		return res, err
	}
	return
}

func (q *Queries) Update(ctx context.Context, company company.Company, id int64) (res company.Company, err error) {
	query := `UPDATE public.company
			  SET name = $1,
			  code = $2,
   			  country = $3,
			  website = $4,
			  phone = $5
			  WHERE id = $6
		     RETURNING id, name, code, country, website, phone`

	row := q.db.QueryRow(ctx, query, company.Name, company.Code, company.Country, company.Website, company.Phone, id)

	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.Code,
		&res.Country,
		&res.Website,
		&res.Phone,
	)
	if err != nil {
		return
	}

	return
}

func (q *Queries) Delete(ctx context.Context, id int64) (int64, error) {
	query := `DELETE FROM public.company WHERE id = $1`

	res, err := q.db.Exec(ctx, query, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (q *Queries) List(ctx context.Context) ([]company.Company, error) {
	query := `SELECT id, name, code, country, website, phone
			  FROM public.company`

	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []company.Company
	for rows.Next() {
		var i company.Company
		if err = rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.Country,
			&i.Website,
			&i.Phone,
		); err != nil {
			return nil, err
		}
		list = append(list, i)
	}
	return list, nil
}

func (q *Queries) ListWithFilter(ctx context.Context, f map[string]string) ([]company.Company, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := builder.Select("id").
		Columns("name").
		Columns("code").
		Columns("country").
		Columns("website").
		Columns("phone").
		From("public.company")

	if f != nil {
		query = postgresql.BuildFilter(query, f)
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := q.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list := make([]company.Company, 0)

	for rows.Next() {
		var i company.Company
		if err = rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.Country,
			&i.Website,
			&i.Phone,
		); err != nil {
			return nil, err
		}
		list = append(list, i)
	}

	return list, nil
}
