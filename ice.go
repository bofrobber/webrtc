package webrtc

import (
	"github.com/pion/sdp/v2"
	"github.com/pion/webrtc/v2/internal/ice"
)

//go:generate go run gen/genaliasdocs.go -- $GOFILE

const (

	// ICETransportPolicyRelay is an alias for ice.ICETransportPolicyRelay
	ICETransportPolicyRelay = ice.ICETransportPolicyRelay

	// ICETransportPolicyAll is an alias for ice.ICETransportPolicyAll
	ICETransportPolicyAll = ice.ICETransportPolicyAll

	// ICETransportStateNew is an alias for ice.ICETransportStateNew
	ICETransportStateNew = ice.ICETransportStateNew

	// ICETransportStateChecking is an alias for ice.ICETransportStateChecking
	ICETransportStateChecking = ice.ICETransportStateChecking

	// ICETransportStateConnected is an alias for ice.ICETransportStateConnected
	ICETransportStateConnected = ice.ICETransportStateConnected

	// ICETransportStateCompleted is an alias for ice.ICETransportStateCompleted
	ICETransportStateCompleted = ice.ICETransportStateCompleted

	// ICETransportStateFailed is an alias for ice.ICETransportStateFailed
	ICETransportStateFailed = ice.ICETransportStateFailed

	// ICETransportStateDisconnected is an alias for ice.ICETransportStateDisconnected
	ICETransportStateDisconnected = ice.ICETransportStateDisconnected

	// ICETransportStateClosed is an alias for ice.ICETransportStateClosed
	ICETransportStateClosed = ice.ICETransportStateClosed

	// ICEConnectionStateNew is an alias for ice.ICEConnectionStateNew
	ICEConnectionStateNew = ice.ICEConnectionStateNew

	// ICEConnectionStateChecking is an alias for ice.ICEConnectionStateChecking
	ICEConnectionStateChecking = ice.ICEConnectionStateChecking

	// ICEConnectionStateConnected is an alias for ice.ICEConnectionStateConnected
	ICEConnectionStateConnected = ice.ICEConnectionStateConnected

	// ICEConnectionStateCompleted is an alias for ice.ICEConnectionStateCompleted
	ICEConnectionStateCompleted = ice.ICEConnectionStateCompleted

	// ICEConnectionStateDisconnected is an alias for ice.ICEConnectionStateDisconnected
	ICEConnectionStateDisconnected = ice.ICEConnectionStateDisconnected

	// ICEConnectionStateFailed is an alias for ice.ICEConnectionStateFailed
	ICEConnectionStateFailed = ice.ICEConnectionStateFailed

	// ICEConnectionStateClosed is an alias for ice.ICEConnectionStateClosed
	ICEConnectionStateClosed = ice.ICEConnectionStateClosed

	// ICEGatheringStateNew is an alias for ice.ICEGatheringStateNew
	ICEGatheringStateNew = ice.ICEGatheringStateNew

	// ICEGatheringStateGathering is an alias for ice.ICEGatheringStateGathering
	ICEGatheringStateGathering = ice.ICEGatheringStateGathering

	// ICEGatheringStateComplete is an alias for ice.ICEGatheringStateComplete
	ICEGatheringStateComplete = ice.ICEGatheringStateComplete

	// ICERoleControlled is an alias for ice.ICERoleControlled
	ICERoleControlled = ice.ICERoleControlled

	// ICERoleControlling is an alias for ice.ICERoleControlling
	ICERoleControlling = ice.ICERoleControlling

	// ICECredentialTypePassword is an alias for ice.ICECredentialTypePassword
	ICECredentialTypePassword = ice.ICECredentialTypePassword

	// ICECredentialTypeOauth is an alias for ice.ICECredentialTypeOauth
	ICECredentialTypeOauth = ice.ICECredentialTypeOauth

	// ICECandidateTypeHost is an alias for ice.ICECandidateTypeHost
	ICECandidateTypeHost = ice.ICECandidateTypeHost

	// ICECandidateTypeSrflx is an alias for ice.ICECandidateTypeSrflx
	ICECandidateTypeSrflx = ice.ICECandidateTypeSrflx

	// ICECandidateTypePrflx is an alias for ice.ICECandidateTypePrflx
	ICECandidateTypePrflx = ice.ICECandidateTypePrflx

	// ICECandidateTypeRelay is an alias for ice.ICECandidateTypeRelay
	ICECandidateTypeRelay = ice.ICECandidateTypeRelay

	// NetworkTypeUDP4 is an alias for ice.NetworkTypeUDP4
	NetworkTypeUDP4 = ice.NetworkTypeUDP4

	// NetworkTypeUDP6 is an alias for ice.NetworkTypeUDP6
	NetworkTypeUDP6 = ice.NetworkTypeUDP6

	// NetworkTypeTCP4 is an alias for ice.NetworkTypeTCP4
	NetworkTypeTCP4 = ice.NetworkTypeTCP4

	// NetworkTypeTCP6 is an alias for ice.NetworkTypeTCP6
	NetworkTypeTCP6 = ice.NetworkTypeTCP6

	// ICEProtocolUDP is an alias for ice.ICEProtocolUDP
	ICEProtocolUDP = ice.ICEProtocolUDP

	// ICEProtocolTCP is an alias for ice.ICEProtocolTCP
	ICEProtocolTCP = ice.ICEProtocolTCP

	// ICEComponentRTP is an alias for ice.ICEComponentRTP
	ICEComponentRTP = ice.ICEComponentRTP

	// ICEComponentRTCP is an alias for ice.ICEComponentRTCP
	ICEComponentRTCP = ice.ICEComponentRTCP

	// ICEGathererStateNew is an alias for ice.ICEGathererStateNew
	ICEGathererStateNew = ice.ICEGathererStateNew

	// ICEGathererStateGathering is an alias for ice.ICEGathererStateGathering
	ICEGathererStateGathering = ice.ICEGathererStateGathering

	// ICEGathererStateComplete is an alias for ice.ICEGathererStateComplete
	ICEGathererStateComplete = ice.ICEGathererStateComplete

	// ICEGathererStateClosed is an alias for ice.ICEGathererStateClosed
	ICEGathererStateClosed = ice.ICEGathererStateClosed
)

type (

	// ICEServer is an alias for ice.ICEServer
	ICEServer = ice.ICEServer

	// ICETransportPolicy is an alias for ice.ICETransportPolicy
	ICETransportPolicy = ice.ICETransportPolicy

	// ICETransport is an alias for ice.ICETransport
	ICETransport = ice.ICETransport

	// ICEGatheringState is an alias for ice.ICEGatheringState
	ICEGatheringState = ice.ICEGatheringState

	// ICEConnectionState is an alias for ice.ICEConnectionState
	ICEConnectionState = ice.ICEConnectionState

	// ICECandidate is an alias for ice.ICECandidate
	ICECandidate = ice.ICECandidate

	// ICEGatherer is an alias for ice.ICEGatherer
	ICEGatherer = ice.ICEGatherer

	// ICEGatherOptions is an alias for ice.ICEGatherOptions
	ICEGatherOptions = ice.ICEGatherOptions

	// ICETransportState is an alias for ice.ICETransportState
	ICETransportState = ice.ICETransportState

	// ICERole is an alias for ice.ICERole
	ICERole = ice.ICERole

	// ICEParameters is an alias for ice.ICEParameters
	ICEParameters = ice.ICEParameters

	// ICECandidateInit is an alias for ice.ICECandidateInit
	ICECandidateInit = ice.ICECandidateInit

	// ICECandidateType is an alias for ice.ICECandidateType
	ICECandidateType = ice.ICECandidateType

	// ICECredentialType is an alias for ice.ICECredentialType
	ICECredentialType = ice.ICECredentialType

	// ICEProtocol is an alias for ice.ICEProtocol
	ICEProtocol = ice.ICEProtocol

	// ICECandidatePair is an alias for ice.ICECandidatePair
	ICECandidatePair = ice.ICECandidatePair

	// ICEComponent is an alias for ice.ICEComponent
	ICEComponent = ice.ICEComponent

	// ICEGathererState is an alias for ice.ICEGathererState
	ICEGathererState = ice.ICEGathererState

	// NetworkType is an alias for ice.NetworkType
	NetworkType = ice.NetworkType

	// OAuthCredential is an alias for ice.OAuthCredential
	OAuthCredential = ice.OAuthCredential
)

var (

	// NewICEGatherer is an alias for ice.NewICEGatherer
	NewICEGatherer = ice.NewICEGatherer

	// NewICETransport is an alias for ice.NewICETransport
	NewICETransport = ice.NewICETransport

	// NewICECandidatePair is an alias for ice.NewICECandidatePair
	NewICECandidatePair = ice.NewICECandidatePair

	// NewICECandidateType is an alias for ice.NewICECandidateType
	NewICECandidateType = ice.NewICECandidateType

	// NewICEProtocol is an alias for ice.NewICEProtocol
	NewICEProtocol = ice.NewICEProtocol

	// ErrNoTurnCredencials is an alias for ice.ErrNoTurnCredencials
	ErrNoTurnCredencials = ice.ErrNoTurnCredencials

	// ErrTurnCredencials is an alias for ice.ErrTurnCredencials
	ErrTurnCredencials = ice.ErrTurnCredencials
)

// NewICEGatherer creates a new NewICEGatherer.
// This constructor is part of the ORTC API. It is not
// meant to be used together with the basic WebRTC API.
func (api *API) NewICEGatherer(opts ICEGatherOptions) (*ICEGatherer, error) {
	return NewICEGatherer(
		api.settingEngine.ephemeralUDP.PortMin,
		api.settingEngine.ephemeralUDP.PortMax,
		api.settingEngine.timeout.ICEConnection,
		api.settingEngine.timeout.ICEKeepalive,
		api.settingEngine.LoggerFactory,
		api.settingEngine.candidates.ICENetworkTypes,
		opts,
	)
}

// NewICETransport creates a new NewICETransport.
// This constructor is part of the ORTC API. It is not
// meant to be used together with the basic WebRTC API.
func (api *API) NewICETransport(gatherer *ICEGatherer) *ICETransport {
	return NewICETransport(gatherer, api.settingEngine.LoggerFactory)
}

// Conversion for package sdp
func newICECandidateFromSDP(c sdp.ICECandidate) (ICECandidate, error) {
	typ, err := NewICECandidateType(c.Typ)
	if err != nil {
		return ICECandidate{}, err
	}
	protocol, err := NewICEProtocol(c.Protocol)
	if err != nil {
		return ICECandidate{}, err
	}
	return ICECandidate{
		Foundation:     c.Foundation,
		Priority:       c.Priority,
		IP:             c.IP,
		Protocol:       protocol,
		Port:           c.Port,
		Component:      c.Component,
		Typ:            typ,
		RelatedAddress: c.RelatedAddress,
		RelatedPort:    c.RelatedPort,
	}, nil
}

func iceCandidateToSDP(c ICECandidate) sdp.ICECandidate {
	return sdp.ICECandidate{
		Foundation:     c.Foundation,
		Priority:       c.Priority,
		IP:             c.IP,
		Protocol:       c.Protocol.String(),
		Port:           c.Port,
		Component:      c.Component,
		Typ:            c.Typ.String(),
		RelatedAddress: c.RelatedAddress,
		RelatedPort:    c.RelatedPort,
	}
}
