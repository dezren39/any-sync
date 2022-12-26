package acllistbuilder

type Key struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Keys struct {
	Derived string `yaml:"Derived"`
	Enc     []*Key `yaml:"Enc"`
	Sign    []*Key `yaml:"Sign"`
	Read    []*Key `yaml:"Read"`
}

type AclChange struct {
	UserAdd *struct {
		Identity          string   `yaml:"identity"`
		EncryptionKey     string   `yaml:"encryptionKey"`
		EncryptedReadKeys []string `yaml:"encryptedReadKeys"`
		Permission        string   `yaml:"permission"`
	} `yaml:"userAdd"`

	UserJoin *struct {
		Identity          string   `yaml:"identity"`
		EncryptionKey     string   `yaml:"encryptionKey"`
		AcceptKey         string   `yaml:"acceptKey"`
		EncryptedReadKeys []string `yaml:"encryptedReadKeys"`
	} `yaml:"userJoin"`

	UserInvite *struct {
		AcceptKey         string   `yaml:"acceptKey"`
		EncryptionKey     string   `yaml:"encryptionKey"`
		EncryptedReadKeys []string `yaml:"encryptedReadKeys"`
		Permissions       string   `yaml:"permissions"`
	} `yaml:"userInvite"`

	UserRemove *struct {
		RemovedIdentity string   `yaml:"removedIdentity"`
		NewReadKey      string   `yaml:"newReadKey"`
		IdentitiesLeft  []string `yaml:"identitiesLeft"`
	} `yaml:"userRemove"`

	UserPermissionChange *struct {
		Identity   string `yaml:"identity"`
		Permission string `yaml:"permission"`
	}
}

type Record struct {
	Identity   string       `yaml:"identity"`
	AclChanges []*AclChange `yaml:"aclChanges"`
	ReadKey    string       `yaml:"readKey"`
}

type Header struct {
	FirstChangeId string `yaml:"firstChangeId"`
	IsWorkspace   bool   `yaml:"isWorkspace"`
}

type Root struct {
	Identity string `yaml:"identity"`
	SpaceId  string `yaml:"spaceId"`
}

type YMLList struct {
	Root    *Root
	Records []*Record `yaml:"records"`

	Keys Keys `yaml:"keys"`
}
