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

func (r *Router) Generate(req GenerateRequest) (string, error) {
	resp, err := r.Primary.Generate(req)
	if err == nil && resp != "" {
		return resp, nil
	}

	if r.Secondary != nil {
		return r.Secondary.Generate(req)
	}

	return "", errors.New("no available model for generation")
}

func (r *Router) Name() string {
	if r.Primary != nil {
		return "Router/" + r.Primary.Name()
	}
	return "Router/no-primary"
}
