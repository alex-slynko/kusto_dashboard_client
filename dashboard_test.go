package kusto_dashboard_client_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/alex-slynko/kusto_dashboard_client"
)

var _ = Describe("Dashboard", func() {

	It("should return proper CreatedAt time", func() {
		dashboard := kusto_dashboard_client.Dashboard{CreatedAt: 1643737646}
		expectedDate := time.Date(2022, 2, 1, 17, 47, 26, 0, time.UTC)
		Expect(dashboard.GetCreatedAt().UTC()).To(Equal(expectedDate))
	})
})
