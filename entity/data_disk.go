package entity

type DataDiskEntity struct {
	Key                 string       `gorm:"unique;type:string;size:64;comment:''" json:"key"`
	Name                string       `gorm:"type:string;size:128;comment:'Data disk name';" json:"name"`
	OwnerUserDataKey    string       `gorm:"type:string;size:64;comment:'Associate the primary key of the owner of the current data disk'" json:"owner_user_data_key"`
	OwnerUser           UserEntity   `gorm:"foreignKey:OwnerUserDataKey;" json:"owner_user"`
	BelongDomainDataKey string       `gorm:"type:string;size:64;comment:'Primary key of the domain to which the data disk belongs.'" json:"belong_domain_data_key"`
	BelongDomain        DomainEntity `gorm:"foreignKey:BelongDomainDataKey;" json:"belong_domain"`
}
