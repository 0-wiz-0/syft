package spdxtagvalue

import (
	"fmt"
	"io"

	"github.com/spdx/tools-golang/tvloader"

	"github.com/anchore/syft/syft/formats/common/spdxhelpers"
	"github.com/anchore/syft/syft/sbom"
)

func decoder(reader io.Reader) (*sbom.SBOM, error) {
	doc, err := tvloader.Load2_3(reader)
	if err != nil {
		return nil, fmt.Errorf("unable to decode spdx-tag-value: %w", err)
	}

	return spdxhelpers.ToSyftModel(doc)
}
