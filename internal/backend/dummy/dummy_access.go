package dummy

// AccessControl ----
type AccessControl struct{}

// Authenticate -----
func (a AccessControl) Authenticate(accessToken string) error { return nil }

// ListTemplates ----
func (a AccessControl) ListTemplates(accessToken string) error { return nil }

// ReadTemplate ----
func (a AccessControl) ReadTemplate(accessToken string, templateName string) error { return nil }

// DeleteTemplate ----
func (a AccessControl) DeleteTemplate(accessToken string, templateName string) error { return nil }

// ManageTemplate ---
func (a AccessControl) ManageTemplate(accessToken string, templateName string) error { return nil }

// CreateTemplate ----
func (a AccessControl) CreateTemplate(accessToken string) error { return nil }

// ListSSHTemplates ----
func (a AccessControl) ListSSHTemplates(accessToken string) error { return nil }

// CreateSSHTemplate ----
func (a AccessControl) CreateSSHTemplate(accessToken string) error { return nil }

// ReadSSHTemplate ----
func (a AccessControl) ReadSSHTemplate(accessToken string, templateName string) error { return nil }

// DeleteSSHTemplate ----
func (a AccessControl) DeleteSSHTemplate(accessToken string, templateName string) error { return nil }

// ManageSSHTemplate ---
func (a AccessControl) ManageSSHTemplate(accessToken string, templateName string) error { return nil }

// CreateSSHCertificate ----
func (a AccessControl) CreateSSHCertificate(accessToken string, templateName string) error {
	return nil
}

// Purge -----------
func (a AccessControl) Purge(accessToken string) error { return nil }

// CRLPurge ----
func (a AccessControl) CRLPurge(accessToken string) error { return nil }

// CreateCertificate -----
func (a AccessControl) CreateCertificate(accessToken string, templateName string) error { return nil }

// RevokeCertificate -----
func (a AccessControl) RevokeCertificate(accessToken string, serialNumber string) error { return nil }

// SignCertificate -----
func (a AccessControl) SignCertificate(accessToken string, templateName string) error { return nil }

// GenerateIntermediateCSR ...
func (a AccessControl) GenerateIntermediateCSR(accessToken string) error { return nil }

// SetIntermediateCertificate ...
func (a AccessControl) SetIntermediateCertificate(accessToken string) error { return nil }

// SetCAChain ...
func (a AccessControl) SetCAChain(accessToken string) error { return nil }
