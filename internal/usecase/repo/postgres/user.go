package postgres

import (
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

// TODO
func (r *Repo) GetUserByNameRepo(userName string) (*entity.User, error) {
	u := entity.User{}
	//sql, args, err := r.Builder.
	//	Select("id, username, email, role_id, phone, password_hash, password_salt, address_id, removed, created_at, updated_at, deleted_at").
	//	From("tbl_users").
	//	Where(squirrel.Eq{"username": userName}).
	//	ToSql()
	//if err != nil {
	//	return u, fmt.Errorf("PostgresRepo - GetUserRepo - r.Builder: %w", err)
	//}
	//ctx := context.Background()
	//row := r.Pool.QueryRow(ctx, sql, args...)
	//if err != nil {
	//	return u, fmt.Errorf("PostgresRepo - GetUserRepo - r.Pool.Query: %w", err)
	//}
	//
	//err = row.Scan(&u.ID, &u.Username, &u.Email, &u.RoleID, &u.Phone, &u.PasswordHash, &u.PasswordSalt, &u.AddressID, &u.Removed, &u.CreateAt, &u.UpdateAt, &u.DeleteAt)
	//if err != nil {
	//	return entity.User{}, fmt.Errorf("UserRepo - GetUsers - rows.Scan: %w", err)
	//}
	//
	//role, err := r.GetRole(u.RoleID)
	//if err != nil {
	//	return entity.User{}, fmt.Errorf("UserRepo - GetUserByNameRepo - GetRoleRepo: %w", err)
	//}
	//u.RoleName = role.Name
	//u.RolePrivileges = role.Privileges
	return &u, nil
}

func (r *Repo) CreateUserRepo(u entity.User) error {
	//err := u.BeforeCreate() // generate PasswordHash and PasswordSalt
	//if err != nil {
	//	return fmt.Errorf("PostgresRepo - CreateUserRepo - u.BeforeCreate(): %w", err)
	//}
	//sql, args, err := r.Builder.
	//	Insert("tbl_users").
	//	Columns("username, email, password_hash, password_salt").
	//	Values(u.Username, u.Email, u.PasswordHash, u.PasswordSalt).
	//	ToSql()
	//if err != nil {
	//	return fmt.Errorf("PostgresRepo - CreateUserRepo - r.Builder: %w", err)
	//}
	//ctx := context.Background()
	//_, err = r.Pool.Exec(ctx, sql, args...)
	//if err != nil {
	//	return fmt.Errorf("PostgresRepo - CreateUserRepo - r.Pool.Exec: %w", err)
	//}

	return nil
}

func (r *Repo) UpdateUserRepo(u entity.User, id int) error {
	//sql, args, err := r.Builder.
	//	Update("tbl_users").
	//	Set("username", u.Username).
	//	Set("email", u.Email).
	//	Set("phone", u.Phone).
	//	Set("address_id", u.AddressID).
	//	Where(squirrel.Eq{"id": id}).
	//	ToSql()
	//if err != nil {
	//	return fmt.Errorf("PostgresRepo - UpdateUserRepo - r.Builder: %w", err)
	//}
	//ctx := context.Background()
	//_, err = r.Pool.Exec(ctx, sql, args...)
	//if err != nil {
	//	return fmt.Errorf("PostgresRepo - UpdateUserRepo - r.Pool.Query: %w", err)
	//}
	return nil
}

func (r *Repo) DeleteUserRepo(id int) error {
	//sql, args, err := r.Builder.
	//	Update("tbl_users").
	//	Set("removed", 1).
	//	Where(squirrel.Eq{"id": id}).
	//	ToSql()
	//if err != nil {
	//	return fmt.Errorf("PostgresRepo - UpdateUserRepo - r.Builder: %w", err)
	//}
	//ctx := context.Background()
	//_, err = r.Pool.Exec(ctx, sql, args...)
	//if err != nil {
	//	return fmt.Errorf("PostgresRepo - UpdateUserRepo - r.Pool.Query: %w", err)
	//}
	return nil
}

func (r *Repo) GetUserRepo(id int) (*entity.User, error) {
	u := entity.User{}
	//sql, args, err := r.Builder.
	//	Select("id, username, email, role_id, phone, address_id, removed, created_at, updated_at, deleted_at").
	//	From("tbl_users").
	//	Where(squirrel.Eq{"id": id}).
	//	ToSql()
	//if err != nil {
	//	return u, fmt.Errorf("PostgresRepo - GetUserRepo - r.Builder: %w", err)
	//}
	//ctx := context.Background()
	//row := r.Pool.QueryRow(ctx, sql, args...)
	//if err != nil {
	//	return u, fmt.Errorf("PostgresRepo - GetUserRepo - r.Pool.Query: %w", err)
	//}
	//
	//err = row.Scan(&u.ID, &u.Username, &u.Email, &u.RoleID, &u.Phone, &u.AddressID, &u.Removed, &u.CreateAt, &u.UpdateAt, &u.DeleteAt)
	//if err != nil {
	//	return entity.User{}, fmt.Errorf("UserRepo - GetUsers - rows.Scan: %w", err)
	//}

	return &u, nil
}

func (r *Repo) GetUsersRepo() (*[]entity.User, error) {
	//sql, _, err := r.Builder.
	//	Select("id, username, email, role_id, phone, address_id, removed, created_at, updated_at, deleted_at").
	//	From("tbl_users").
	//	ToSql()
	//if err != nil {
	//	return nil, fmt.Errorf("UserRepo - GetUsers - r.Builder: %w", err)
	//}
	//ctx := context.Background()         //!!!
	//rows, err := r.Pool.Query(ctx, sql) //!!!
	//if err != nil {
	//	return nil, fmt.Errorf("UserRepo - GetUsers - r.Pool.Query: %w", err)
	//}
	//defer rows.Close()

	entities := make([]entity.User, 0, _defaultEntityCap)

	//for rows.Next() {
	//	u := entity.User{}
	//	err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.RoleID, &u.Phone, &u.AddressID, &u.Removed, &u.CreateAt, &u.UpdateAt, &u.DeleteAt)
	//	if err != nil {
	//		return nil, fmt.Errorf("UserRepo - GetUsers - rows.Scan: %w", err)
	//	}
	//	entities = append(entities, u)
	//}

	return &entities, nil
}
