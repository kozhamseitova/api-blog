package pgrepo

//func (p *Postgres) CreateArticle(ctx context.Context, a *entity.Article) error {
//	u := entity.User{}
//	query := fmt.Sprintf(`
//			INSERT INTO %s (
//			                title, -- 1
//			                description, -- 2
//			                user_id, -- 3
//			                password -- 4
//			                )
//			VALUES ($1, $2, $3, $4)
//			`, articlesTable)
//
//	_, err := p.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.Password)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (p *Postgres) UpdateArticle(ctx context.Context, a *entity.Article) error {
//	return nil
//}
//
//func (p *Postgres) DeleteArticleByID(ctx context.Context, id int64) error {
//	return nil
//}
//
//func (p *Postgres) GetArticleByID(ctx context.Context, id int64) (*entity.Article, error) {
//	return nil, nil
//}
//
//func (p *Postgres) GetAllArticles(ctx context.Context) ([]entity.Article, error) {
//	return nil, nil
//}
//
//func (p *Postgres) GetArticlesByUserID(ctx context.Context, userID int64) ([]entity.Article, error) {
//	return nil, nil
//}
