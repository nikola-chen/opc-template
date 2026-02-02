package ai

import "errors"

type Router struct {
	Primary   Client
	Secondary Client
}

func (r *Router) Fix(req FixRequest) (FixResponse, error) {
	resp, err := r.Primary.Fix(req)
	if err == nil && resp.Patch != "" {
		return resp, nil
	}

	if r.Secondary != nil {
		return r.Secondary.Fix(req)
	}

	return FixResponse{}, errors.New("no available model")

}
