package api

import (
	"context"
	"encoding/json"
	"errors"
	"journey/internal/api/spec"
	"journey/internal/pgstore"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"github.com/go-playground/validator/v10"
)

type store interface {
	CreateTrip(context.Context, *pgxpool.Pool, spec.CreateTripRequest) (uuid.UUID, error)
	GetParticipant( context.Context, uuid.UUID) (pgstore.Participant, error)
	ConfirmParticipant(context.Context,  uuid.UUID) (error)
}

type mailer interface {
	SendConfirmTripEmailToTripOwner(uuid.UUID) error
}

type API struct {
	store 		store
	logger 		*zap.Logger
	validator 	*validator.Validate
	pool 		*pgxpool.Pool
	mailer 		mailer

}

func NewAPI(pool *pgxpool.Pool, logger *zap.Logger, mailer mailer) API {
	validator := validator.New(validator.WithRequiredStructEnabled())

	return API {
		pgstore.New(pool),
		logger,
		validator,
		pool,
		mailer,
	}
}

// Confirms a participant on a trip.
// (PATCH /participants/{participantId}/confirm)
func (api API) PatchParticipantsParticipantIDConfirm(
	w http.ResponseWriter,
	r *http.Request, 
	participantID string,
	) *spec.Response {

	id, err := uuid.Parse(participantID)
	if err != nil {
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(
			spec.Error{Message: "UUID Inválido"},
		)
	}
	participant, err := api.store.GetParticipant(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return spec.PatchParticipantsParticipantIDConfirmJSON400Response(
				spec.Error{Message: "Participante não encontrado"},
			)
		}
		api.logger.Error("Failed to get participant", zap.Error(err), zap.String("Participant ID", participantID))
		return spec.GetTripsTripIDActivitiesJSON400Response(
			spec.Error{Message: "Something went wrong. Try again later."},
		)
	}
	if participant.IsConfirmed {
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(
			spec.Error{Message: "Participante já confirmado"},
		)
	}
	if err := api.store.ConfirmParticipant(r.Context(), id); err != nil {
		api.logger.Error(
			"Failed to confirm participant", 
			zap.Error(err), 
			zap.String("Participant ID", participantID),
			)
		return spec.PostTripsTripIDActivitiesJSON400Response(
			spec.Error{Message: "Something went wrong. Try again later."},
		)
	}
	return spec.PatchParticipantsParticipantIDConfirmJSON204Response(nil)
}

// Create a new trip
// (POST /trips)
func (api API) PostTrips(w http.ResponseWriter, r *http.Request) *spec.Response {
	var body spec.CreateTripRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return spec.PostTripsJSON400Response(spec.Error{
			Message: "Invalid JSON: " + err.Error(),
		})
	}

	if err := api.validator.Struct(body); err != nil {
		return spec.PostTripsJSON400Response(spec.Error{
			Message: "Invalid input: " + err.Error(),
		})
	}
	tripID, err := api.store.CreateTrip(
		r.Context(), 
		api.pool, 
		body,
	);	if err != nil {
		return spec.PostTripsJSON400Response(
			spec.Error{Message: "Failed to create trip, try again later"},
		)
	}

	go func () {
		if err := api.mailer.SendConfirmTripEmailToTripOwner(tripID); err != nil {
			api.logger.Error("Failed to send email on PostTrips", 
			zap.Error(err),
			zap.String("trip_id", tripID.String()))
		}
	} ()
	return spec.PostTripsJSON201Response(spec.CreateTripResponse{TripID: tripID.String()})
}	

// Get a trip details.
// (GET /trips/{tripId})
func (api API) GetTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Update a trip.
// (PUT /trips/{tripId})
func (api API) PutTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip activities.
// (GET /trips/{tripId}/activities)
func (api API) GetTripsTripIDActivities(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Create a trip activity.
// (POST /trips/{tripId}/activities)
func (api API) PostTripsTripIDActivities(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Confirm a trip and send e-mail invitations.
// (GET /trips/{tripId}/confirm)
func (api API) GetTripsTripIDConfirm(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Invite someone to the trip.
// (POST /trips/{tripId}/invites)
func (api API) PostTripsTripIDInvites(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip links.
// (GET /trips/{tripId}/links)
func (api API) GetTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Create a trip link.
// (POST /trips/{tripId}/links)
func (api API) PostTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip participants.
// (GET /trips/{tripId}/participants)
func (api API) GetTripsTripIDParticipants(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

