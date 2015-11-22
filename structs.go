package discordgo

import (
	"net"

	"golang.org/x/net/websocket"
)

// TODO: Eventually everything here gets moved to a better place.

// A Message stores all data related to a specific Discord message.
type Message struct {
	ID              string       `json:"id"`
	Author          User         `json:"author"`
	Content         string       `json:"content"`
	Attachments     []Attachment `json:"attachments"`
	Tts             bool         `json:"tts"`
	Embeds          []Embed      `json:"embeds"`
	Timestamp       string       `json:"timestamp"`
	MentionEveryone bool         `json:"mention_everyone"`
	EditedTimestamp string       `json:"edited_timestamp"`
	Mentions        []User       `json:"mentions"`
	ChannelID       string       `json:"channel_id"`
}

// An Attachment stores data for message attachments.
type Attachment struct { //TODO figure this out
}

// An Embed stores data for message embeds.
type Embed struct { // TODO figure this out
}

// A VoiceRegion stores data for a specific voice region server.
type VoiceRegion struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Hostname string `json:"sample_hostname"`
	Port     int    `json:"sample_port"`
}

// A VoiceICE stores data for voice ICE servers.
type VoiceICE struct {
	TTL     string      `json:"ttl"`
	Servers []ICEServer `json:"servers"`
}

// A ICEServer stores data for a specific voice ICE server.
type ICEServer struct {
	URL        string `json:"url"`
	Username   string `json:"username"`
	Credential string `json:"credential"`
}

// A Invite stores all data related to a specific Discord Guild or Channel invite.
type Invite struct {
	MaxAge    int     `json:"max_age"`
	Code      string  `json:"code"`
	Guild     Guild   `json:"guild"`
	Revoked   bool    `json:"revoked"`
	CreatedAt string  `json:"created_at"` // TODO make timestamp
	Temporary bool    `json:"temporary"`
	Uses      int     `json:"uses"`
	MaxUses   int     `json:"max_uses"`
	Inviter   User    `json:"inviter"`
	XkcdPass  bool    `json:"xkcdpass"`
	Channel   Channel `json:"channel"`
}

// A Channel holds all data related to an individual Discord channel.
type Channel struct {
	ID                   string                `json:"id"`
	GuildID              string                `json:"guild_id"`
	Name                 string                `json:"name"`
	Topic                string                `json:"topic"`
	Position             int                   `json:"position"`
	Type                 string                `json:"type"`
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
	IsPrivate            bool                  `json:"is_private"`
	LastMessageID        string                `json:"last_message_id"`
	Recipient            User                  `json:"recipient"`
}

// A PermissionOverwrite holds permission overwrite data for a Channel
type PermissionOverwrite struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Deny  int    `json:"deny"`
	Allow int    `json:"allow"`
}

// A Guild holds all data related to a specific Discord Guild.  Guilds are also
// sometimes referred to as Servers in the Discord client.
type Guild struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	Icon           string       `json:"icon"`
	Region         string       `json:"region"`
	AfkTimeout     int          `json:"afk_timeout"`
	AfkChannelID   string       `json:"afk_channel_id"`
	EmbedChannelID string       `json:"embed_channel_id"`
	EmbedEnabled   bool         `json:"embed_enabled"`
	OwnerID        string       `json:"owner_id"`
	Large          bool         `json:"large"`     // ??
	JoinedAt       string       `json:"joined_at"` // make this a timestamp
	Roles          []Role       `json:"roles"`
	Members        []Member     `json:"members"`
	Presences      []Presence   `json:"presences"`
	Channels       []Channel    `json:"channels"`
	VoiceStates    []VoiceState `json:"voice_states"`
}

// A Role stores information about Discord guild member roles.
type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Managed     bool   `json:"managed"`
	Color       int    `json:"color"`
	Hoist       bool   `json:"hoist"`
	Position    int    `json:"position"`
	Permissions int    `json:"permissions"`
}

// A VoiceState stores the voice states of Guilds
type VoiceState struct {
	UserID    string `json:"user_id"`
	Suppress  bool   `json:"suppress"`
	SessionID string `json:"session_id"`
	SelfMute  bool   `json:"self_mute"`
	SelfDeaf  bool   `json:"self_deaf"`
	Mute      bool   `json:"mute"`
	Deaf      bool   `json:"deaf"`
	ChannelID string `json:"channel_id"`
}

// A Presence stores the online, offline, or idle and game status of Guild members.
type Presence struct {
	User   User   `json:"user"`
	Status string `json:"status"`
	GameID int    `json:"game_id"`
}

// A Member stores user information for Guild members.
type Member struct {
	GuildID  string   `json:"guild_id"`
	JoinedAt string   `json:"joined_at"`
	Deaf     bool     `json:"deaf"`
	Mute     bool     `json:"mute"`
	User     User     `json:"user"`
	Roles    []string `json:"roles"`
}

// A Session represents a connection to the Discord REST API.
// token : The authentication token returned from Discord
// Debug : If set to ture debug logging will be displayed.
type Session struct {
	Token     string // Authentication token for this session
	Debug     bool   // Debug for printing JSON request/responses
	Cache     int    // number in X to cache some responses
	SessionID string // from websocket READY packet

	// Settable Callback functions for Websocket Events
	OnEvent                   func(*Session, Event) // should Event be *Event?
	OnReady                   func(*Session, Ready)
	OnTypingStart             func(*Session, TypingStart)
	OnMessageCreate           func(*Session, Message)
	OnMessageUpdate           func(*Session, Message)
	OnMessageDelete           func(*Session, MessageDelete)
	OnMessageAck              func(*Session, MessageAck)
	OnPresenceUpdate          func(*Session, PresenceUpdate)
	OnVoiceStateUpdate        func(*Session, VoiceState)
	OnChannelCreate           func(*Session, Channel)
	OnChannelUpdate           func(*Session, Channel)
	OnChannelDelete           func(*Session, Channel)
	OnGuildCreate             func(*Session, Guild)
	OnGuildUpdate             func(*Session, Guild)
	OnGuildDelete             func(*Session, Guild)
	OnGuildMemberAdd          func(*Session, Member)
	OnGuildMemberRemove       func(*Session, Member)
	OnGuildMemberDelete       func(*Session, Member) // which is it?
	OnGuildMemberUpdate       func(*Session, Member)
	OnGuildRoleCreate         func(*Session, GuildRole)
	OnGuildRoleUpdate         func(*Session, GuildRole)
	OnGuildRoleDelete         func(*Session, GuildRoleDelete)
	OnGuildIntegrationsUpdate func(*Session, GuildIntegrationsUpdate)

	wsConn *websocket.Conn
	//TODO, add bools for like.
	// are we connnected to websocket?
	// have we authenticated to login?
	// lets put all the general session
	// tracking and infos here.. clearly

	// Everything below here is used for Voice testing.
	// This stuff is almost guarenteed to change a lot
	// and is even a bit hackish right now.
	VwsConn    *websocket.Conn // new for voice
	VSessionID string
	VToken     string
	VEndpoint  string
	VGuildID   string
	VChannelID string
	Vop2       VoiceOP2
	UDPConn    *net.UDPConn
}

// A User stores all data for an individual Discord user.
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"Avatar"`
	Verified bool   `json:"verified"`
	//Discriminator int    `json:"discriminator,string"` // TODO: See below
}

// Discriminator sometimes comes as a string
// and sometimes it comes as a int.  Weird.
// to avoid errors I've just commented it out
// but it doesn't seem to just kill the whole
// program.  Heartbeat is taken on READY even
// with error and the system continues to read
// it just doesn't seem able to handle this one
// field correctly.  Need to research this more.

// A PrivateChannel stores all data for a specific user private channel.
type PrivateChannel struct {
	ID            string `json:"id"`
	IsPrivate     bool   `json:"is_private"`
	LastMessageID string `json:"last_message_id"`
	Recipient     User   `json:"recipient"`
} // merge with channel?

// A Settings stores data for a specific users Discord client settings.
type Settings struct {
	RenderEmbeds          bool     `json:"render_embeds"`
	InlineEmbedMedia      bool     `json:"inline_embed_media"`
	EnableTtsCommand      bool     `json:"enable_tts_command"`
	MessageDisplayCompact bool     `json:"message_display_compact"`
	Locale                string   `json:"locale"`
	ShowCurrentGame       bool     `json:"show_current_game"`
	Theme                 string   `json:"theme"`
	MutedChannels         []string `json:"muted_channels"`
}

// An Event provides a basic initial struct for all websocket event.
type Event struct {
	Type      string          `json:"t"`
	State     int             `json:"s"`
	Operation int             `json:"o"`
	Direction int             `json:"dir"`
	RawData   json.RawMessage `json:"d"`
}

// A Ready stores all data for the websocket READY event.
type Ready struct {
	Version           int           `json:"v"`
	SessionID         string        `json:"session_id"`
	HeartbeatInterval time.Duration `json:"heartbeat_interval"`
	User              User          `json:"user"`
	ReadState         []ReadState
	PrivateChannels   []PrivateChannel
	Guilds            []Guild
}

// ReadState might need to move? Gives me the read status
// of all my channels when first connecting. I think :)

// A ReadState stores data on the read state of channels.
type ReadState struct {
	MentionCount  int
	LastMessageID string `json:"last_message_id"`
	ID            string `json:"id"`
}

// A TypingStart stores data for the typing start websocket event.
type TypingStart struct {
	UserID    string `json:"user_id"`
	ChannelID string `json:"channel_id"`
	Timestamp int    `json:"timestamp"`
}

// A PresenceUpdate stores data for the pressence update websocket event.
type PresenceUpdate struct {
	User    User     `json:"user"`
	Status  string   `json:"status"`
	Roles   []string `json:"roles"`
	GuildID string   `json:"guild_id"`
	GameID  int      `json:"game_id"`
}

// A MessageAck stores data for the message ack websocket event.
type MessageAck struct {
	MessageID string `json:"message_id"`
	ChannelID string `json:"channel_id"`
}

// A MessageDelete stores data for the message delete websocket event.
type MessageDelete struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
} // so much like MessageAck..

// A GuildIntegrationsUpdate stores data for the guild integrations update
// websocket event.
type GuildIntegrationsUpdate struct {
	GuildID string `json:"guild_id"`
}

// A GuildRole stores data for guild role websocket events.
type GuildRole struct {
	Role    Role   `json:"role"`
	GuildID string `json:"guild_id"`
}

// A GuildRoleDelete stores data for the guild role delete websocket event.
type GuildRoleDelete struct {
	RoleID  string `json:"role_id"`
	GuildID string `json:"guild_id"`
}