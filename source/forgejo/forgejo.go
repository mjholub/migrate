package forgejo

import (
	"code.gitea.io/sdk/gitea"
	"fmt"
	"github.com/golang-migrate/migrate/v4/source"
	nurl "net/url"
)

const (
  ErrNoUserInfo = "no username:token provided"
  ErrNoAccessToken = "no access token"
)

func init() {
	source.Register("forgejo", &ForgejoGitea{})
}

type ForgejoGitea struct {
	client *gitea.Client
	url    string

	projectID   string
	path        string
	treeOptions *treeOpts
	mivgrations *source.Migrations
}

type treeOpts struct {
	user string
	repo string
	// recursive bool
}

func (fg *ForgejoGitea) Open(url string) (source.Driver, error) {
	u, err := nurl.Parse(url)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}

  if u.User == nil {
    return nil, fmt.Errorf(ErrNoUserInfo) 
  }

  password, ok := u.User.Password()
  if !ok {
    return nil, fmt.Errorf(ErrNoAccessToken)
  }

  fgn := &ForgejoGitea{
    client: gitea.NewClient(u.String(), nil),
    url: url,
    migrations: source.NewMigrations(),
}

	if u.Host != "" {

repo, res, err := fgn.client.GetRepo(u.User, u.Path)
if err != nil {
      return nil, fmt.Errorf("error getting repo: %w", err)
  }
}
  
	pe := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(pe) < 1 {
		return nil, ErrInvalidProjectID
	}

