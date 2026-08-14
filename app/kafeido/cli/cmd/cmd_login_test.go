package cmd

import "testing"

// The oauth2 login command's initial KafeidoServiceAppLogin call must never
// attach a token from ~/.kafeidoconfig: the user is not authenticated yet,
// and a stale or expired token in that file makes AuthN reject the login
// request itself before the login handler ever runs - "Permission denied...
// haven't login first" on a request whose whole purpose is to log in
// (grandturks#273). NewBasicLoginCommand already gets this right by passing
// a literal nil for the same call.
func TestOauth2LoginRequestAuthInformerNeverAttachesAToken(t *testing.T) {
	if got := oauth2LoginRequestAuthInformer(); got != nil {
		t.Errorf("oauth2LoginRequestAuthInformer() = %v, want nil - "+
			"the login request must not carry a saved token", got)
	}
}
