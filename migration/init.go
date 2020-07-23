package migration

import "todo_back/models"

type DBInit struct{}

func (i DBInit) RollUp() {
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}, &models.Task{})
}

func (i DBInit) RollBack() {
	DB.DropTable("users", "tasks")
}

func (i DBInit) GetName() string {
	return "DBInit"
}
