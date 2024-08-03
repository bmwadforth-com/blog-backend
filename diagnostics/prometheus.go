package diagnostics

import "github.com/prometheus/client_golang/prometheus"

var (
	ArticlesCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "articles_count",
			Help: "Total number of times articles has been called",
		},
		[]string{"article_slug"},
	)
)
