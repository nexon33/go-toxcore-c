package tox

/*
#include <tox/tox.h>
#include <tox/toxav.h>
*/
import "C"

const (
	PUBLIC_KEY_SIZE           = int(C.TOX_PUBLIC_KEY_SIZE)
	SECRET_KEY_SIZE           = int(C.TOX_SECRET_KEY_SIZE)
	ADDRESS_SIZE              = int(C.TOX_ADDRESS_SIZE)
	MAX_NAME_LENGTH           = int(C.TOX_MAX_NAME_LENGTH)
	MAX_STATUS_MESSAGE_LENGTH = int(C.TOX_MAX_STATUS_MESSAGE_LENGTH)
	MAX_FRIEND_REQUEST_LENGTH = int(C.TOX_MAX_FRIEND_REQUEST_LENGTH)
	MAX_MESSAGE_LENGTH        = int(C.TOX_MAX_MESSAGE_LENGTH)
	MAX_CUSTOM_PACKET_SIZE    = int(C.TOX_MAX_CUSTOM_PACKET_SIZE)
	HASH_LENGTH               = int(C.TOX_HASH_LENGTH)
	FILE_ID_LENGTH            = int(C.TOX_FILE_ID_LENGTH)
	MAX_FILENAME_LENGTH       = int(C.TOX_MAX_FILENAME_LENGTH)
)

type UserStatus int

const (
	USER_STATUS_NONE = int(C.TOX_USER_STATUS_NONE)
	USER_STATUS_AWAY = int(C.TOX_USER_STATUS_AWAY)
	USER_STATUS_BUSY = int(C.TOX_USER_STATUS_BUSY)
)

type ConnectionType int

const (
	CONNECTION_NONE = int(C.TOX_CONNECTION_NONE)
	CONNECTION_TCP  = int(C.TOX_CONNECTION_TCP)
	CONNECTION_UDP  = int(C.TOX_CONNECTION_UDP)
)

type FileControlType int

const (
	FILE_CONTROL_RESUME = int(C.TOX_FILE_CONTROL_RESUME)
	FILE_CONTROL_PAUSE  = int(C.TOX_FILE_CONTROL_PAUSE)
	FILE_CONTROL_CANCEL = int(C.TOX_FILE_CONTROL_CANCEL)
)

type FileKind uint32

const (
	FILE_KIND_DATA   = uint32(C.TOX_FILE_KIND_DATA)
	FILE_KIND_AVATAR = uint32(C.TOX_FILE_KIND_AVATAR)
)

type GroupchatType int

const (
	GROUPCHAT_TYPE_TEXT  = uint8(C.TOX_CONFERENCE_TYPE_TEXT)
	GROUPCHAT_TYPE_AV    = uint8(C.TOX_CONFERENCE_TYPE_AV)
	CONFERENCE_TYPE_TEXT = int(C.TOX_CONFERENCE_TYPE_TEXT)
	CONFERENCE_TYPE_AV   = int(C.TOX_CONFERENCE_TYPE_AV)
)

type ChatChangeType int

const (
	CHAT_CHANGE_PEER_ADD                     = uint8(C.TOX_CONFERENCE_STATE_CHANGE_PEER_JOIN)
	CHAT_CHANGE_PEER_DEL                     = uint8(C.TOX_CONFERENCE_STATE_CHANGE_PEER_EXIT)
	CHAT_CHANGE_PEER_NAME                    = uint8(C.TOX_CONFERENCE_STATE_CHANGE_PEER_NAME_CHANGE)
	CONFERENCE_STATE_CHANGE_PEER_JOIN        = int(C.TOX_CONFERENCE_STATE_CHANGE_PEER_JOIN)
	CONFERENCE_STATE_CHANGE_PEER_EXIT        = int(C.TOX_CONFERENCE_STATE_CHANGE_PEER_EXIT)
	CONFERENCE_STATE_CHANGE_PEER_NAME_CHANGE = int(C.TOX_CONFERENCE_STATE_CHANGE_PEER_NAME_CHANGE)
)

type CallControlType int

const (
	FRIEND_CALL_STATE_ERROR       = int(C.TOXAV_FRIEND_CALL_STATE_ERROR)
	FRIEND_CALL_STATE_FINISHED    = int(C.TOXAV_FRIEND_CALL_STATE_FINISHED)
	FRIEND_CALL_STATE_SENDING_A   = int(C.TOXAV_FRIEND_CALL_STATE_SENDING_A)
	FRIEND_CALL_STATE_SENDING_V   = int(C.TOXAV_FRIEND_CALL_STATE_SENDING_V)
	FRIEND_CALL_STATE_ACCEPTING_A = int(C.TOXAV_FRIEND_CALL_STATE_ACCEPTING_A)
	FRIEND_CALL_STATE_ACCEPTING_V = int(C.TOXAV_FRIEND_CALL_STATE_ACCEPTING_V)
)

type MessageType int

const (
	MESSAGE_TYPE_NORMAL = int(C.TOX_MESSAGE_TYPE_NORMAL)
	MESSAGE_TYPE_ACTION = int(C.TOX_MESSAGE_TYPE_ACTION)
)
