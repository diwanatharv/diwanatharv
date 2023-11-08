package models

import "time"

type Organization struct {
	Id                   uint      `json:"id" gorm:"primary_key"`
	OrganizationName     string    `json:"organization_name"`
	AdminEmail           string    `json:"admin_email"`
	SiteURL              string    `json:"site_url"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Status               string    `json:"status"`
	AuthenticationMethod string    `json:"authentication_method"`
}
type LoginCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}
type Tenant struct {
	Id             uint      `json:"id" gorm:"primary_key"`
	TenantName     string    `json:"tenant_name"`
	AdminEmail     string    `json:"admin_email"`
	SiteURL        string    `json:"site_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	OrganizationId int       `json:"organization_id"`
	Status         string    `json:"status"`
}

type User struct {
	UserId           int     `gorm:"column:user_id;primary_key"`
	EmailAddress     string  `gorm:"column:email_address"`
	PhoneNumber      string  `gorm:"column:phone_number"`
	LogOnName        string  `gorm:"column:logon_name"`
	City             string  `gorm:"column:city"`
	Country          string  `gorm:"column:country"`
	Industry         string  `gorm:"column:industry"`
	Organization     string  `gorm:"column:organization"`
	CompanyHeadcount string  `gorm:"column:company_headcount"`
	FirstName        string  `gorm:"column:firstname"`
	LastName         string  `gorm:"column:lastname"`
	Address          string  `gorm:"column:address"`
	Password         string  `gorm:"column:user_password"`
	DomainId         []uint8 `gorm:"column:domain_id"`
	Status           string  `gorm:"column:status"`
	OtpMethod        string  `gorm:"column:otp_method"`
	Metadata         string  `gorm:"column:metadata"`
	Dn               string  `gorm:"column:dn"`
	UserRoleID       int     `gorm:"column:user_role_id"`
	OrgID            int     `gorm:"column:org_id"`
}

func (User) TableName() string {
	return "users"
}

// CREATE TABLE did.epm_machines (
// 	machine_id int8 NOT NULL DEFAULT nextval('did.epm_machines_seq'::regclass),
// 	machine_key text NOT NULL,
// 	public_ip_address varchar(512) NOT NULL,
// 	auth_type varchar(255) NOT NULL,
// 	private_ip_address varchar(255) NOT NULL,
// 	os_id text NOT NULL,
// 	status varchar(255) NOT NULL DEFAULT ''::character varying,
// 	hostname varchar(255) NOT NULL DEFAULT ''::character varying,
// 	ip_address varchar(255) NOT NULL DEFAULT ''::character varying,
// 	localuser varchar(255) NOT NULL DEFAULT ''::character varying,
// 	password_policy_id int4 NULL DEFAULT 0,
// 	auth_code varchar(512) NOT NULL DEFAULT ''::character varying,
// 	jump_server_id int4 NULL DEFAULT 0,
// 	domain_id int4 NOT NULL,
// 	vnc_password varchar(255) NOT NULL DEFAULT ''::character varying,
// 	factors varchar(512) NULL,
// 	status_guacd bool NULL,
// 	status_services bool NULL,
// 	instance_id int8 NOT NULL DEFAULT 0,
// 	CONSTRAINT epm_machines_pkey PRIMARY KEY (machine_id)
// );

type EpmMachine struct {
	MachineId        int64  `gorm:"column:machine_id;primary_key"`
	MachineKey       string `gorm:"column:machine_key"`
	PublicIpAddress  string `gorm:"column:public_ip_address"`
	AuthType         string `gorm:"column:auth_type"`
	PrivateIpAddress string `gorm:"column:private_ip_address"`
	OsId             string `gorm:"column:os_id"`
	Status           string `gorm:"column:status"`
	Hostname         string `gorm:"column:hostname"`
	IpAddress        string `gorm:"column:ip_address"`
	Localuser        string `gorm:"column:localuser"`
	PasswordPolicyId int    `gorm:"column:password_policy_id"`
	AuthCode         string `gorm:"column:auth_code"`
	JumpServerId     int    `gorm:"column:jump_server_id"`
	DomainId         int    `gorm:"column:domain_id"`
	VncPassword      string `gorm:"column:vnc_password"`
	Factors          string `gorm:"column:factors"`
	StatusGuacd      bool   `gorm:"column:status_guacd"`
	StatusServices   bool   `gorm:"column:status_services"`
	InstanceId       int64  `gorm:"column:instance_id"`
}

type GetEndpointListRequest struct {
	Email string `json:"email" validate:"required,email"`
}
