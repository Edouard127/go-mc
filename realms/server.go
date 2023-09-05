package realms

import (
	"errors"
	"fmt"
)

type Server struct {
	ID                   int      `json:"id"`
	RemoteSubscriptionID string   `json:"remoteSubscriptionId"`
	Owner                string   `json:"owner"`
	OwnerUUID            string   `json:"ownerUUID"`
	Name                 string   `json:"name"`
	MOTD                 string   `json:"motd"`
	State                string   `json:"state"`
	DaysLeft             int      `json:"daysLeft"`
	Expired              bool     `json:"expired"`
	ExpiredTrial         bool     `json:"expiredTrial"`
	WorldType            string   `json:"worldType"`
	Players              []string `json:"players"`
	MaxPlayers           int      `json:"maxPlayers"`
	MiniGameName         *string  `json:"minigameName,omitempty"`
	MiniGameID           *int     `json:"minigameId,omitempty"`
	MinigameImage        *string  `json:"minigameImage,omitempty"`
	ActiveSlot           int      `json:"activeSlot"`
	//Slots                interface{}
	Member bool `json:"member"`
}

type Backup struct {
	ID               int   `json:"backupId"`
	LastModifiedDate int64 `json:"lastModifiedDate"`
	Size             int64 `json:"size"`
	Metadata         struct {
		Difficulty   string `json:"game_difficulty"`
		Name         string `json:"name"`
		Version      string `json:"version"`
		EnabledPacks struct {
			RessourcePacks []string `json:"resourcePacks"`
			BehaviorPacks  []string `json:"behaviorPacks"`
		}
		Description string `json:"description"`
		Mode        string `json:"game_mode"`
		Type        string `json:"world_type"`
	} `json:"metadata"`
}

// Worlds return a list of servers that the user is invited to or owns.
func (r *Realms) Worlds() ([]Server, error) {
	var resp struct {
		Servers []Server
		*Error
	}

	err := r.get("/worlds", &resp)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		err = resp.Error
	}

	return resp.Servers, err
}

// Server returns a single server listing about a server.
// you must be the owner of the server.
func (r *Realms) Server(ID int) (s Server, err error) {
	var resp = struct {
		*Server
		*Error
	}{Server: &s}

	err = r.get(fmt.Sprintf("/worlds/%d", ID), &resp)
	if err != nil {
		return
	}

	if resp.Error != nil {
		err = resp.Error
	}

	return
}

// Address used to get the IP address for a server.
// Call TOS before you call this function.
func (r *Realms) Address(s Server) (string, error) {
	var resp struct {
		Address       string
		PendingUpdate bool

		ResourcePackUrl  *string
		ResourcePackHash *string

		*Error
	}

	err := r.get(fmt.Sprintf("/worlds/v1/%d/join/pc", s.ID), &resp)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		err = resp.Error
		return "", err
	}

	if resp.PendingUpdate {
		return "", errors.New("pending update")
	}
	return resp.Address, err
}

// Backups returns a list of backups for the world.
func (r *Realms) Backups(s Server) ([]Backup, error) {
	var bs []Backup
	err := r.get(fmt.Sprintf("/worlds/%d/backups", s.ID), &bs)

	return bs, err
}

func (r *Realms) Download() (link, resURL, resHash string) {
	var resp struct {
		DownloadLink     string `json:"downloadLink"`
		ResourcePackURL  string `json:"resourcePackUrl,omitempty"`
		ResourcePackHash string `json:"resourcePackHash,omitempty"`
		*Error
	}
	// TODO: What is the ID?
	if r.get("/worlds/$ID/slot/1/download", &resp) != nil {
		return "", "", ""
	}

	return resp.DownloadLink, resp.ResourcePackURL, resp.ResourcePackHash
}

// Ops returns a list of operators for this server.
// You must own this server to view this.
func (r *Realms) Ops(s Server) (ops []string, err error) {
	err = r.get(fmt.Sprintf("/ops/%d", s.ID), &ops)
	return
}

// SubscriptionLife returns the current life of a server subscription.
func (r *Realms) SubscriptionLife(s Server) (startDate int64, daysLeft int, Type string, err error) {
	var resp = struct {
		StartDate        *int64
		DaysLeft         *int
		SubscriptionType *string
	}{
		StartDate:        &startDate,
		DaysLeft:         &daysLeft,
		SubscriptionType: &Type,
	}

	err = r.get(fmt.Sprintf("/subscriptions/%d", s.ID), &resp)
	return
}
