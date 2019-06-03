// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cruisecontrol

import (
	"fmt"
	"strings"

	"github.com/banzaicloud/kafka-operator/pkg/resources/kafka"
	"github.com/banzaicloud/kafka-operator/pkg/resources/templates"
	"github.com/banzaicloud/kafka-operator/pkg/util"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"

	banzaicloudv1alpha1 "github.com/banzaicloud/kafka-operator/pkg/apis/banzaicloud/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

func (r *Reconciler) configMap(log logr.Logger) runtime.Object {
	configMap := &corev1.ConfigMap{
		ObjectMeta: templates.ObjectMeta(configAndVolumeName, labelSelector, r.KafkaCluster),
		Data: map[string]string{
			"cruisecontrol.properties": fmt.Sprintf(`
    # Copyright 2017 LinkedIn Corp. Licensed under the BSD 2-Clause License (the "License"). See License in the project root for license information.
    #
    # This is an example property file for Kafka Cruise Control. See KafkaCruiseControlConfig for more details.
    # Configuration for the metadata client.
    # =======================================
    # The Kafka cluster to control.
    bootstrap.servers=%s:%d
    # The maximum interval in milliseconds between two metadata refreshes.
    #metadata.max.age.ms=300000
    # Client id for the Cruise Control. It is used for the metadata client.
    #client.id=kafka-cruise-control
    # The size of TCP send buffer bytes for the metadata client.
    #send.buffer.bytes=131072
    # The size of TCP receive buffer size for the metadata client.
    #receive.buffer.bytes=131072
    # The time to wait before disconnect an idle TCP connection.
    #connections.max.idle.ms=540000
    # The time to wait before reconnect to a given host.
    #reconnect.backoff.ms=50
    # The time to wait for a response from a host after sending a request.
    #request.timeout.ms=30000
    # Configurations for the load monitor
    # =======================================
    # The number of metric fetcher thread to fetch metrics for the Kafka cluster
    num.metric.fetchers=1
    # The metric sampler class
    metric.sampler.class=com.linkedin.kafka.cruisecontrol.monitor.sampling.CruiseControlMetricsReporterSampler
    # Configurations for CruiseControlMetricsReporterSampler
    metric.reporter.topic.pattern=__CruiseControlMetrics
    # The sample store class name
    sample.store.class=com.linkedin.kafka.cruisecontrol.monitor.sampling.KafkaSampleStore
    # The config for the Kafka sample store to save the partition metric samples
    partition.metric.sample.store.topic=__KafkaCruiseControlPartitionMetricSamples
    # The config for the Kafka sample store to save the model training samples
    broker.metric.sample.store.topic=__KafkaCruiseControlModelTrainingSamples
    # The replication factor of Kafka metric sample store topic
    sample.store.topic.replication.factor=2
    # The config for the number of Kafka sample store consumer threads
    num.sample.loading.threads=8
    # The partition assignor class for the metric samplers
    metric.sampler.partition.assignor.class=com.linkedin.kafka.cruisecontrol.monitor.sampling.DefaultMetricSamplerPartitionAssignor
    # The metric sampling interval in milliseconds
    metric.sampling.interval.ms=120000
    # The partition metrics window size in milliseconds
    partition.metrics.window.ms=300000
    # The number of partition metric windows to keep in memory
    num.partition.metrics.windows=1
    # The minimum partition metric samples required for a partition in each window
    min.samples.per.partition.metrics.window=1
    # The broker metrics window size in milliseconds
    broker.metrics.window.ms=300000
    # The number of broker metric windows to keep in memory
    num.broker.metrics.windows=20
    # The minimum broker metric samples required for a partition in each window
    min.samples.per.broker.metrics.window=1
    # The configuration for the BrokerCapacityConfigFileResolver (supports JBOD and non-JBOD broker capacities)
    capacity.config.file=config/capacity.json
    #capacity.config.file=config/capacityJBOD.json
    # Configurations for the analyzer
    # =======================================
    # The list of goals to optimize the Kafka cluster for with pre-computed proposals
    default.goals=com.linkedin.kafka.cruisecontrol.analyzer.goals.ReplicaCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.DiskCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkInboundCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkOutboundCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.CpuCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.ReplicaDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.PotentialNwOutGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.DiskUsageDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkInboundUsageDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkOutboundUsageDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.CpuUsageDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.TopicReplicaDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.LeaderBytesInDistributionGoal
    # The list of supported goals
    goals=com.linkedin.kafka.cruisecontrol.analyzer.goals.ReplicaCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.DiskCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkInboundCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkOutboundCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.CpuCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.ReplicaDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.PotentialNwOutGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.DiskUsageDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkInboundUsageDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkOutboundUsageDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.CpuUsageDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.TopicReplicaDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.LeaderBytesInDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.kafkaassigner.KafkaAssignerDiskUsageDistributionGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.PreferredLeaderElectionGoal
    # The list of supported hard goals
    hard.goals=com.linkedin.kafka.cruisecontrol.analyzer.goals.ReplicaCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.DiskCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkInboundCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkOutboundCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.CpuCapacityGoal
    # The minimum percentage of well monitored partitions out of all the partitions
    min.monitored.partition.percentage=0.95
    # The balance threshold for CPU
    cpu.balance.threshold=1.1
    # The balance threshold for disk
    disk.balance.threshold=1.1
    # The balance threshold for network inbound utilization
    network.inbound.balance.threshold=1.1
    # The balance threshold for network outbound utilization
    network.outbound.balance.threshold=1.1
    # The balance threshold for the replica count
    replica.count.balance.threshold=1.1
    # The capacity threshold for CPU in percentage
    cpu.capacity.threshold=0.8
    # The capacity threshold for disk in percentage
    disk.capacity.threshold=0.8
    # The capacity threshold for network inbound utilization in percentage
    network.inbound.capacity.threshold=0.8
    # The capacity threshold for network outbound utilization in percentage
    network.outbound.capacity.threshold=0.8
    # The threshold to define the cluster to be in a low CPU utilization state
    cpu.low.utilization.threshold=0.0
    # The threshold to define the cluster to be in a low disk utilization state
    disk.low.utilization.threshold=0.0
    # The threshold to define the cluster to be in a low network inbound utilization state
    network.inbound.low.utilization.threshold=0.0
    # The threshold to define the cluster to be in a low disk utilization state
    network.outbound.low.utilization.threshold=0.0
    # The metric anomaly percentile upper threshold
    metric.anomaly.percentile.upper.threshold=90.0
    # The metric anomaly percentile lower threshold
    metric.anomaly.percentile.lower.threshold=10.0
    # How often should the cached proposal be expired and recalculated if necessary
    proposal.expiration.ms=60000
    # The maximum number of replicas that can reside on a broker at any given time.
    max.replicas.per.broker=10000
    # The number of threads to use for proposal candidate precomputing.
    num.proposal.precompute.threads=1
    # the topics that should be excluded from the partition movement.
    #topics.excluded.from.partition.movement
    # Configurations for the executor
    # =======================================
    # The zookeeper connect of the Kafka cluster
    zookeeper.connect=%s/
    # The max number of partitions to move in/out on a given broker at a given time.
    num.concurrent.partition.movements.per.broker=10
    # The interval between two execution progress checks.
    execution.progress.check.interval.ms=10000
    # Configurations for anomaly detector
    # =======================================
    # The goal violation notifier class
    anomaly.notifier.class=com.linkedin.kafka.cruisecontrol.detector.notifier.SelfHealingNotifier
    # The metric anomaly finder class
    metric.anomaly.finder.class=com.linkedin.kafka.cruisecontrol.detector.KafkaMetricAnomalyFinder
    # The anomaly detection interval
    anomaly.detection.interval.ms=10000
    # The goal violation to detect.
    anomaly.detection.goals=com.linkedin.kafka.cruisecontrol.analyzer.goals.ReplicaCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.DiskCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkInboundCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.NetworkOutboundCapacityGoal,com.linkedin.kafka.cruisecontrol.analyzer.goals.CpuCapacityGoal
    # The interested metrics for metric anomaly analyzer.
    metric.anomaly.analyzer.metrics=BROKER_PRODUCE_LOCAL_TIME_MS_MAX,BROKER_PRODUCE_LOCAL_TIME_MS_MEAN,BROKER_CONSUMER_FETCH_LOCAL_TIME_MS_MAX,BROKER_CONSUMER_FETCH_LOCAL_TIME_MS_MEAN,BROKER_FOLLOWER_FETCH_LOCAL_TIME_MS_MAX,BROKER_FOLLOWER_FETCH_LOCAL_TIME_MS_MEAN,BROKER_LOG_FLUSH_TIME_MS_MAX,BROKER_LOG_FLUSH_TIME_MS_MEAN
    ## Adjust accordingly if your metrics reporter is an older version and does not produce these metrics.
    #metric.anomaly.analyzer.metrics=BROKER_PRODUCE_LOCAL_TIME_MS_50TH,BROKER_PRODUCE_LOCAL_TIME_MS_999TH,BROKER_CONSUMER_FETCH_LOCAL_TIME_MS_50TH,BROKER_CONSUMER_FETCH_LOCAL_TIME_MS_999TH,BROKER_FOLLOWER_FETCH_LOCAL_TIME_MS_50TH,BROKER_FOLLOWER_FETCH_LOCAL_TIME_MS_999TH,BROKER_LOG_FLUSH_TIME_MS_50TH,BROKER_LOG_FLUSH_TIME_MS_999TH
    # The zk path to store failed broker information.
    failed.brokers.zk.path=/CruiseControlBrokerList
    # Topic config provider class
    topic.config.provider.class=com.linkedin.kafka.cruisecontrol.config.KafkaTopicConfigProvider
    # The cluster configurations for the KafkaTopicConfigProvider
    cluster.configs.file=config/clusterConfigs.json
    # The maximum time in milliseconds to store the response and access details of a completed user task.
    completed.user.task.retention.time.ms=21600000
    # The maximum time in milliseconds to retain the demotion history of brokers.
    demotion.history.retention.time.ms=86400000
    # The maximum number of completed user tasks for which the response and access details will be cached.
    max.cached.completed.user.tasks=100
    # The maximum number of user tasks for concurrently running in async endpoints across all users.
    max.active.user.tasks=5
    # Enable self healing for all anomaly detectors, unless the particular anomaly detector is explicitly disabled
    self.healing.enabled=true
    # Enable self healing for broker failure detector
    #self.healing.broker.failure.enabled=true
    # Enable self healing for goal violation detector
    #self.healing.goal.violation.enabled=true
    # Enable self healing for metric anomaly detector
    #self.healing.metric.anomaly.enabled=true
    # configurations for the webserver
    # ================================
    # HTTP listen port
    webserver.http.port=9090
    # HTTP listen address
    webserver.http.address=0.0.0.0
    # Whether CORS support is enabled for API or not
    webserver.http.cors.enabled=false
    # Value for Access-Control-Allow-Origin
    webserver.http.cors.origin=http://localhost:8080/
    # Value for Access-Control-Request-Method
    webserver.http.cors.allowmethods=OPTIONS,GET,POST
    # Headers that should be exposed to the Browser (Webapp)
    # This is a special header that is used by the
    # User Tasks subsystem and should be explicitly
    # Enabled when CORS mode is used as part of the
    # Admin Interface
    webserver.http.cors.exposeheaders=User-Task-ID
    # REST API default prefix
    # (dont forget the ending *)
    webserver.api.urlprefix=/kafkacruisecontrol/*
    # Location where the Cruise Control frontend is deployed
    webserver.ui.diskpath=./cruise-control-ui/dist/
    # URL path prefix for UI
    # (dont forget the ending *)
    webserver.ui.urlprefix=/*
    # Time After which request is converted to Async
    webserver.request.maxBlockTimeMs=10000
    # Default Session Expiry Period
    webserver.session.maxExpiryTimeMs=60000
    # Session cookie path
    webserver.session.path=/
    # Server Access Logs
    webserver.accesslog.enabled=true
    # Location of HTTP Request Logs
    webserver.accesslog.path=access.log
    # HTTP Request Log retention days
    webserver.accesslog.retention.days=14
`, fmt.Sprintf(kafka.HeadlessServiceTemplate, r.KafkaCluster.Name), r.KafkaCluster.Spec.ListenersConfig.InternalListeners[0].ContainerPort, strings.Join(r.KafkaCluster.Spec.ZKAddresses, ",")) +
				generateSSLConfig(&r.KafkaCluster.Spec.ListenersConfig),
			"capacity.json": `
{
      "brokerCapacities":[
        {
          "brokerId": "-1",
          "capacity": {
            "DISK": "10000",
            "CPU": "100",
            "NW_IN": "10000",
            "NW_OUT": "10000"
          },
          "doc": "This is the default capacity. Capacity unit used for disk is in MB, cpu is in percentage, network throughput is in KB."
        }
      ]
    }`,
			"clusterConfigs.json": `
{
      "min.insync.replicas": 2
    }
`,
			"log4j.properties": `
log4j.rootLogger = INFO, FILE
    log4j.appender.FILE=org.apache.log4j.FileAppender
    log4j.appender.FILE.File=/dev/stdout
    log4j.appender.FILE.layout=org.apache.log4j.PatternLayout
    log4j.appender.FILE.layout.conversionPattern=%-6r [%15.15t] %-5p %30.30c %x - %m%n
`,
			"log4j2.xml": `
<?xml version="1.0" encoding="UTF-8"?>
    <Configuration status="INFO">
        <Appenders>
            <File name="Console" fileName="/dev/stdout">
                <PatternLayout pattern="%d{yyy-MM-dd HH:mm:ss.SSS} [%t] %-5level %logger{36} - %msg%n"/>
            </File>
        </Appenders>
        <Loggers>
            <Root level="info">
                <AppenderRef ref="Console" />
            </Root>
        </Loggers>
    </Configuration>
`,
		},
	}
	return configMap
}

func generateSSLConfig(l *banzaicloudv1alpha1.ListenersConfig) (res string) {
	if l.SSLSecrets != nil && util.IsSSLEnabledForInternalCommunication(l.InternalListeners) {
		res = `
security.protocol=SSL
ssl.truststore.location=/var/run/secrets/java.io/keystores/client.truststore.jks
ssl.keystore.location=/var/run/secrets/java.io/keystores/client.keystore.jks
`
	}
	return
}
