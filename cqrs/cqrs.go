package cqrs

import (
	"fmt"
	"time"

	"talentpitch/utils"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Modelos Dinámicos para Comandos y Eventos

type Command map[string]interface{}
type Event map[string]interface{}

// Función para crear un nuevo comando inmutable
func NewCommand(commandType string, data map[string]interface{}) Command {
	return Command{
		"type": commandType,
		"data": data,
	}
}

// Función para crear un nuevo evento inmutable
func NewEvent(eventType string, data map[string]interface{}) Event {
	return Event{
		"type": eventType,
		"data": data,
		"time": time.Now(),
		"id":   uuid.New(),
	}
}

// Validador de Esquemas Dinámicos
var validate = validator.New()

func BuildCommand(input map[string]interface{}) (Event, error) {
	switch input["type"] {
	case "RegisterUser":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":   "required,uuid",
			"UserName": "required,alpha",
			"UserType": "required,oneof=Talent Sponsor Representative",
			"Email":    "required,email",
			"Password": "required,min=8",
		})
		if err != nil {
			// Manejo de errores de validación
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return UserRegistered(input["UserId"].(string), input["UserName"].(string), input["UserType"].(string)), nil

	case "CreateTalentSearch":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":         "required,uuid",
			"SearchCriteria": "required,min=3,max=100",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("TalentSearchCreated", input), nil

	case "UploadTalentVideo":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":           "required,uuid",
			"VideoId":          "required,uuid",
			"VideoTitle":       "required,min=3,max=100",
			"VideoDescription": "required,min=10,max=1000",
			"Skills":           "required,dive,alpha,min=1",
			"Categories":       "required,dive,alpha,min=1",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("TalentVideoUploaded", input), nil

	case "CreateCallForTalentVideo":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":           "required,uuid",
			"CallId":           "required,uuid",
			"CallCriteria":     "required,min=3,max=1000",
			"Deadline":         "required,datetime",
			"VideoTitle":       "required,min=3,max=100",
			"VideoDescription": "required,min=10,max=1000",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("CallForTalentVideoCreated", input), nil

	case "SubmitVideoForCall":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":  "required,uuid",
			"VideoId": "required,uuid",
			"CallId":  "required,uuid",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("VideoSubmittedForCall", input), nil

	case "SelectFeaturedVideos":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":           "required,uuid",
			"CallId":           "required,uuid",
			"SelectedVideoIds": "required,dive,uuid",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("VideosSelectedForCall", input), nil

	case "CreateTalentPlaylist":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":       "required,uuid",
			"PlaylistId":   "required,uuid",
			"PlaylistName": "required,min=3,max=100",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("TalentPlaylistCreated", input), nil

	case "AddVideoToPlaylist":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":     "required,uuid",
			"PlaylistId": "required,uuid",
			"VideoId":    "required,uuid",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("VideoAddedToPlaylist", input), nil

	case "RemoveVideoFromPlaylist":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":     "required,uuid",
			"PlaylistId": "required,uuid",
			"VideoId":    "required,uuid",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("VideoRemovedFromPlaylist", input), nil

	case "CreateMatch":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":       "required,uuid",
			"TargetUserId": "required,uuid",
			"MatchType":    "required,oneof=Contract Collaborate Sponsor Train Mediate",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("MatchCreated", input), nil

	case "ProposeCollaboration":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":          "required,uuid",
			"TargetUserId":    "required,uuid",
			"ProposalDetails": "required,min=10,max=1000",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("CollaborationProposed", input), nil

	case "RespondToCollaborationProposal":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":     "required,uuid",
			"ProposalId": "required,uuid",
			"Response":   "required,oneof=Accept Reject",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		if input["Response"].(string) == "Accept" {
			return NewEvent("CollaborationProposalAccepted", input), nil
		}
		return NewEvent("CollaborationProposalRejected", input), nil

	case "DeleteUserAccount":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId": "required,uuid",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("UserAccountDeleted", input), nil

	case "UnmatchUsers":
		err := validate.ValidateMap(input, map[string]interface{}{
			"UserId":       "required,uuid",
			"TargetUserId": "required,uuid",
		})
		if err != nil {
			return nil, fmt.Errorf("Errors: %s", utils.ReduceErrorsToString(err))
		}
		return NewEvent("UsersUnmatched", input), nil

	default:
		return nil, fmt.Errorf("unknown command type: %s", input["type"])
	}
}

// Comandos

func RegisterUser(userId, userName, userType, email, password string) Command {
	return NewCommand("RegisterUser", map[string]interface{}{
		"UserId":   userId,
		"UserName": userName,
		"UserType": userType,
		"Email":    email,
		"Password": password,
	})
}

func CreateTalentSearch(userId, searchCriteria string) Command {
	return NewCommand("CreateTalentSearch", map[string]interface{}{
		"UserId":         userId,
		"SearchCriteria": searchCriteria,
	})
}

func UploadTalentVideo(userId, videoId, videoTitle, videoDescription string, skills, categories []string) Command {
	return NewCommand("UploadTalentVideo", map[string]interface{}{
		"UserId":           userId,
		"VideoId":          videoId,
		"VideoTitle":       videoTitle,
		"VideoDescription": videoDescription,
		"Skills":           skills,
		"Categories":       categories,
	})
}

func CreateCallForTalentVideo(userId, callId, callCriteria, deadline, videoTitle, videoDescription string) Command {
	return NewCommand("CreateCallForTalentVideo", map[string]interface{}{
		"UserId":           userId,
		"CallId":           callId,
		"CallCriteria":     callCriteria,
		"Deadline":         deadline,
		"VideoTitle":       videoTitle,
		"VideoDescription": videoDescription,
	})
}

func SubmitVideoForCall(userId, videoId, callId string) Command {
	return NewCommand("SubmitVideoForCall", map[string]interface{}{
		"UserId":  userId,
		"VideoId": videoId,
		"CallId":  callId,
	})
}

func SelectFeaturedVideos(userId, callId string, selectedVideoIds []string) Command {
	return NewCommand("SelectFeaturedVideos", map[string]interface{}{
		"UserId":           userId,
		"CallId":           callId,
		"SelectedVideoIds": selectedVideoIds,
	})
}

func CreateTalentPlaylist(userId, playlistId, playlistName string) Command {
	return NewCommand("CreateTalentPlaylist", map[string]interface{}{
		"UserId":       userId,
		"PlaylistId":   playlistId,
		"PlaylistName": playlistName,
	})
}

func AddVideoToPlaylist(userId, playlistId, videoId string) Command {
	return NewCommand("AddVideoToPlaylist", map[string]interface{}{
		"UserId":     userId,
		"PlaylistId": playlistId,
		"VideoId":    videoId,
	})
}

func RemoveVideoFromPlaylist(userId, playlistId, videoId string) Command {
	return NewCommand("RemoveVideoFromPlaylist", map[string]interface{}{
		"UserId":     userId,
		"PlaylistId": playlistId,
		"VideoId":    videoId,
	})
}

func CreateMatch(userId, targetUserId, matchType string) Command {
	return NewCommand("CreateMatch", map[string]interface{}{
		"UserId":       userId,
		"TargetUserId": targetUserId,
		"MatchType":    matchType,
	})
}

func ProposeCollaboration(userId, targetUserId, proposalDetails string) Command {
	return NewCommand("ProposeCollaboration", map[string]interface{}{
		"UserId":          userId,
		"TargetUserId":    targetUserId,
		"ProposalDetails": proposalDetails,
	})
}

func RespondToCollaborationProposal(userId, proposalId, response string) Command {
	return NewCommand("RespondToCollaborationProposal", map[string]interface{}{
		"UserId":     userId,
		"ProposalId": proposalId,
		"Response":   response,
	})
}

func DeleteUserAccount(userId string) Command {
	return NewCommand("DeleteUserAccount", map[string]interface{}{
		"UserId": userId,
	})
}

func UnmatchUsers(userId, targetUserId string) Command {
	return NewCommand("UnmatchUsers", map[string]interface{}{
		"UserId":       userId,
		"TargetUserId": targetUserId,
	})
}

// Eventos

func UserRegistered(userId, userName, userType string) Event {
	return NewEvent("UserRegistered", map[string]interface{}{
		"UserId":   userId,
		"UserName": userName,
		"UserType": userType,
	})
}

func TalentSearchCreated(searchId, userId, searchCriteria string) Event {
	return NewEvent("TalentSearchCreated", map[string]interface{}{
		"SearchId":       searchId,
		"UserId":         userId,
		"SearchCriteria": searchCriteria,
	})
}

func TalentVideoUploaded(videoId, userId, videoTitle, videoDescription string, skills, categories []string) Event {
	return NewEvent("TalentVideoUploaded", map[string]interface{}{
		"VideoId":          videoId,
		"UserId":           userId,
		"VideoTitle":       videoTitle,
		"VideoDescription": videoDescription,
		"Skills":           skills,
		"Categories":       categories,
	})
}

func CallForTalentVideoCreated(callId, userId, callCriteria, deadline, videoTitle, videoDescription string) Event {
	return NewEvent("CallForTalentVideoCreated", map[string]interface{}{
		"CallId":           callId,
		"UserId":           userId,
		"CallCriteria":     callCriteria,
		"Deadline":         deadline,
		"VideoTitle":       videoTitle,
		"VideoDescription": videoDescription,
	})
}

func VideoSubmittedForCall(submissionId, userId, videoId, callId string) Event {
	return NewEvent("VideoSubmittedForCall", map[string]interface{}{
		"SubmissionId": submissionId,
		"UserId":       userId,
		"VideoId":      videoId,
		"CallId":       callId,
	})
}

func VideosSelectedForCall(callId, userId string, selectedVideoIds []string) Event {
	return NewEvent("VideosSelectedForCall", map[string]interface{}{
		"CallId":           callId,
		"UserId":           userId,
		"SelectedVideoIds": selectedVideoIds,
	})
}

func CallWinnerSelected(callId, userId, winnerVideoId string) Event {
	return NewEvent("CallWinnerSelected", map[string]interface{}{
		"CallId":        callId,
		"UserId":        userId,
		"WinnerVideoId": winnerVideoId,
	})
}

func TalentPlaylistCreated(playlistId, userId, playlistName string) Event {
	return NewEvent("TalentPlaylistCreated", map[string]interface{}{
		"PlaylistId":   playlistId,
		"UserId":       userId,
		"PlaylistName": playlistName,
	})
}

func VideoAddedToPlaylist(playlistId, userId, videoId string) Event {
	return NewEvent("VideoAddedToPlaylist", map[string]interface{}{
		"PlaylistId": playlistId,
		"UserId":     userId,
		"VideoId":    videoId,
	})
}

func VideoRemovedFromPlaylist(playlistId, userId, videoId string) Event {
	return NewEvent("VideoRemovedFromPlaylist", map[string]interface{}{
		"PlaylistId": playlistId,
		"UserId":     userId,
		"VideoId":    videoId,
	})
}

func MatchCreated(matchId, userId, targetUserId, matchType string) Event {
	return NewEvent("MatchCreated", map[string]interface{}{
		"MatchId":      matchId,
		"UserId":       userId,
		"TargetUserId": targetUserId,
		"MatchType":    matchType,
	})
}

func UsersUnmatched(userId, targetUserId string) Event {
	return NewEvent("UsersUnmatched", map[string]interface{}{
		"UserId":       userId,
		"TargetUserId": targetUserId,
	})
}

func CollaborationProposed(proposalId, userId, targetUserId, proposalDetails string) Event {
	return NewEvent("CollaborationProposed", map[string]interface{}{
		"ProposalId":      proposalId,
		"UserId":          userId,
		"TargetUserId":    targetUserId,
		"ProposalDetails": proposalDetails,
	})
}

func CollaborationProposalAccepted(proposalId, userId, targetUserId string) Event {
	return NewEvent("CollaborationProposalAccepted", map[string]interface{}{
		"ProposalId":   proposalId,
		"UserId":       userId,
		"TargetUserId": targetUserId,
	})
}

func CollaborationProposalRejected(proposalId, userId, targetUserId string) Event {
	return NewEvent("CollaborationProposalRejected", map[string]interface{}{
		"ProposalId":   proposalId,
		"UserId":       userId,
		"TargetUserId": targetUserId,
	})
}

func UserAccountDeleted(userId string) Event {
	return NewEvent("UserAccountDeleted", map[string]interface{}{
		"UserId": userId,
	})
}

func UserEventsArchived(userId string, archivedEventIds []string) Event {
	return NewEvent("UserEventsArchived", map[string]interface{}{
		"UserId":           userId,
		"ArchivedEventIds": archivedEventIds,
	})
}
