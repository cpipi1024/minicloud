package db

// 资源抽象
type Resource struct {
	Name         string `gorm:"column:name;type:varchar(100)" json:"name"`           //文件名
	Size         uint   `gorm:"column:size;type:varchar(100)" json:"size"`           //文件大小
	Extension    string `gorm:"column:extension;type:varchar(100)" json:"extension"` //扩展类型
	Mime         string `gorm:"column:mime;type:varchar(100)" json:"mime"`           //媒体类型
	Path         string `gorm:"column:path;type:varchar(300)" json:"path"`           //相对路径
	IsDir        bool   `gorm:"column:is_dir" json:"is_dir"`                         //是否文件夹
	LastModified uint64 `gorm:"column:last_modified" json:"last_modified"`           //最后修改时间 unix时间戳
}
