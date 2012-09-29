package model

import (
    "container/list"
    "database/sql"
)

type User struct {
    id uint64
    username, email string
    used_space uint64
}

type UserModel struct {
    db *sql.DB
}

type Users *list.List

func (model *UserModel) GetAll(uid int) (Users, error) {
    //db, err := sql.Open("mymysql", "tcp:localhost:3306*yunio2/root/root")
    rows, err := model.db.Query("SELECT ? FROM user_files LIMIT 10", "*")
    if err != nil {
        return nil, err
    }

    var lst = list.New()
    var path, name string
    //loop to get all data
    for rows.Next() {
        err = rows.Scan(&path, &name)
        if err != nil {
            return nil, err
        }
        lst.PushBack(User{}
        fmt.Println(path, name)
    }
    return lst
}
