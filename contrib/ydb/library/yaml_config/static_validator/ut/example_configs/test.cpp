#include <contrib/ydb/library/yaml_config/static_validator/builders.h>

#include <library/cpp/testing/unittest/registar.h>
#include <contrib/ydb/library/yaml_config/validator/validator.h>
#include <contrib/ydb/library/yaml_config/validator/validator_builder.h>
#include <contrib/ydb/library/yaml_config/static_validator/builders.h>

using namespace NYamlConfig::NValidator;
using TIssue = TValidationResult::TIssue;

bool HasOnlyThisIssues(TValidationResult result, TVector<TIssue> issues) {
    if (result.Issues.size() != issues.size()) {
        Cerr << "Issue counts are differend. List of actul issues:" << Endl;
        Cerr << result;
        Cerr << "------------- List of Expected Issues: " << Endl;
        Cerr << TValidationResult(issues);
        Cerr << "------------- End of issue List" << Endl;
        return false;
    }
    Sort(result.Issues);
    Sort(issues);
    for (size_t i = 0; i < issues.size(); ++i) {
        if (result.Issues[i] != issues[i]) {
            Cerr << "Issues are differend. List of actul issues:" << Endl;
            Cerr << result;
            Cerr << "------------- List of Expected Issues: " << Endl;
            Cerr << TValidationResult(issues);
            Cerr << "------------- End of issue List" << Endl;
            return false;
        }
    }
    return true;
}

bool Valid(TValidationResult result) {
    if (result.Ok()) return true;

    Cerr << "List of issues:" << Endl;
    Cerr << result;
    Cerr << "------------- End of issue list: " << Endl;
    return false;
}

Y_UNIT_TEST_SUITE(StaticConfigExamples) {
    Y_UNIT_TEST(SingleNodeWithFile) {
        auto v = StaticConfigBuilder().CreateValidator();

        const char* yaml =
        "static_erasure: none\n"
        "host_configs:\n"
        "- drive:\n"
        "  - path: /tmp/pdisk.data\n"
        "    type: SSD\n"
        "  host_config_id: 1\n"
        "hosts:\n"
        "- host: localhost\n"
        "  host_config_id: 1\n"
        "  port: 19001\n"
        "  walle_location:\n"
        "    body: 1\n"
        "    data_center: '1'\n"
        "    rack: '1'\n"
        "domains_config:\n"
        "  domain:\n"
        "  - name: Root\n"
        "    storage_pool_types:\n"
        "    - kind: ssd\n"
        "      pool_config:\n"
        "        box_id: 1\n"
        "        erasure_species: none\n"
        "        kind: ssd\n"
        "        pdisk_filter:\n"
        "        - property:\n"
        "          - type: SSD\n"
        "        vdisk_kind: Default\n"
        "  state_storage:\n"
        "  - ring:\n"
        "      node:\n"
        "      - 1\n"
        "      nto_select: 1\n"
        "    ssid: 1\n"
        "table_service_config:\n"
        "  sql_version: 1\n"
        "actor_system_config:\n"
        "  executor:\n"
        "  - name: System\n"
        "    spin_threshold: 0\n"
        "    threads: 2\n"
        "    type: BASIC\n"
        "  - name: User\n"
        "    spin_threshold: 0\n"
        "    threads: 3\n"
        "    type: BASIC\n"
        "  - name: Batch\n"
        "    spin_threshold: 0\n"
        "    threads: 2\n"
        "    type: BASIC\n"
        "  - name: IO\n"
        "    threads: 1\n"
        "    time_per_mailbox_micro_secs: 100\n"
        "    type: IO\n"
        "  - name: IC\n"
        "    spin_threshold: 10\n"
        "    threads: 1\n"
        "    time_per_mailbox_micro_secs: 100\n"
        "    type: BASIC\n"
        "  scheduler:\n"
        "    progress_threshold: 10000\n"
        "    resolution: 256\n"
        "    spin_threshold: 0\n"
        "blob_storage_config:\n"
        "  service_set:\n"
        "    groups:\n"
        "    - erasure_species: none\n"
        "      rings:\n"
        "      - fail_domains:\n"
        "        - vdisk_locations:\n"
        "          - node_id: 1\n"
        "            path: /tmp/pdisk.data\n"
        "            pdisk_category: SSD\n"
        "channel_profile_config:\n"
        "  profile:\n"
        "  - channel:\n"
        "    - erasure_species: none\n"
        "      pdisk_category: 0\n"
        "      storage_pool_kind: ssd\n"
        "    - erasure_species: none\n"
        "      pdisk_category: 0\n"
        "      storage_pool_kind: ssd\n"
        "    - erasure_species: none\n"
        "      pdisk_category: 0\n"
        "      storage_pool_kind: ssd\n"
        "    profile_id: 0\n";

        Y_ENSURE(Valid(v.Validate(yaml)));
    }

    Y_UNIT_TEST(BLOCK42) {
        auto v = StaticConfigBuilder().CreateValidator();
        const char* yaml = 
        "static_erasure: block-4-2\n"
        "host_configs:\n"
        "- drive:\n"
        "    - path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "      type: SSD\n"
        "    - path: /dev/disk/by-partlabel/ydb_disk_ssd_02\n"
        "      type: SSD\n"
        "  host_config_id: 1\n"
        "hosts:\n"
        "- host: ydb-node-zone-a-1.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 1\n"
        "    data_center: \"zone-a\"\n"
        "    rack: \"1\"\n"
        "- host: ydb-node-zone-a-2.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 2\n"
        "    data_center: \"zone-a\"\n"
        "    rack: \"2\"\n"
        "- host: ydb-node-zone-a-3.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 3\n"
        "    data_center: \"zone-a\"\n"
        "    rack: \"3\"\n"
        "\n"
        "- host: ydb-node-zone-a-4.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 4\n"
        "    data_center: \"zone-a\"\n"
        "    rack: \"4\"\n"
        "- host: ydb-node-zone-a-5.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 5\n"
        "    data_center: \"zone-a\"\n"
        "    rack: \"5\"\n"
        "- host: ydb-node-zone-a-6.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 6\n"
        "    data_center: \"zone-a\"\n"
        "    rack: \"6\"\n"
        "\n"
        "- host: ydb-node-zone-a-7.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 7\n"
        "    data_center: \"zone-a\"\n"
        "    rack: \"7\"\n"
        "- host: ydb-node-zone-a-8.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 8\n"
        "    data_center: \"zone-a\"\n"
        "    rack: \"8\"\n"
        "\n"
        "domains_config:\n"
        "  domain:\n"
        "    - name: Root\n"
        "      storage_pool_types:\n"
        "        - kind: ssd\n"
        "          pool_config:\n"
        "            box_id: 1\n"
        "            erasure_species: block-4-2\n"
        "            kind: ssd\n"
        "            pdisk_filter:\n"
        "            - property:\n"
        "                - type: SSD\n"
        "            vdisk_kind: Default\n"
        "  state_storage:\n"
        "    - ring:\n"
        "        node: [1, 2, 3, 4, 5, 6, 7, 8]\n"
        "        nto_select: 5\n"
        "      ssid: 1\n"
        "table_service_config:\n"
        "  sql_version: 1\n"
        "actor_system_config:\n"
        "  executor:\n"
        "    - name: System\n"
        "      threads: 2\n"
        "      type: BASIC\n"
        "    - name: User\n"
        "      threads: 3\n"
        "      type: BASIC\n"
        "    - name: Batch\n"
        "      threads: 2\n"
        "      type: BASIC\n"
        "    - name: IO\n"
        "      threads: 1\n"
        "      time_per_mailbox_micro_secs: 100\n"
        "      type: IO\n"
        "    - name: IC\n"
        "      spin_threshold: 10\n"
        "      threads: 1\n"
        "      time_per_mailbox_micro_secs: 100\n"
        "      type: BASIC\n"
        "  scheduler:\n"
        "    progress_threshold: 10000\n"
        "    resolution: 256\n"
        "    spin_threshold: 0\n"
        "blob_storage_config:\n"
        "  service_set:\n"
        "    groups:\n"
        "    - erasure_species: block-4-2\n"
        "      rings:\n"
        "        - fail_domains:\n"
        "            - vdisk_locations:\n"
        "                - node_id: \"ydb-node-zone-a-1.local\"\n"
        "                  pdisk_category: SSD\n"
        "                  path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "            - vdisk_locations:\n"
        "                - node_id: \"ydb-node-zone-a-2.local\"\n"
        "                  pdisk_category: SSD\n"
        "                  path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "            - vdisk_locations:\n"
        "                - node_id: \"ydb-node-zone-a-3.local\"\n"
        "                  pdisk_category: SSD\n"
        "                  path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "            - vdisk_locations:\n"
        "                - node_id: \"ydb-node-zone-a-4.local\"\n"
        "                  pdisk_category: SSD\n"
        "                  path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "            - vdisk_locations:\n"
        "                - node_id: \"ydb-node-zone-a-5.local\"\n"
        "                  pdisk_category: SSD\n"
        "                  path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "            - vdisk_locations:\n"
        "                - node_id: \"ydb-node-zone-a-6.local\"\n"
        "                  pdisk_category: SSD\n"
        "                  path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "            - vdisk_locations:\n"
        "                - node_id: \"ydb-node-zone-a-7.local\"\n"
        "                  pdisk_category: SSD\n"
        "                  path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "            - vdisk_locations:\n"
        "                - node_id: \"ydb-node-zone-a-8.local\"\n"
        "                  pdisk_category: SSD\n"
        "                  path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "channel_profile_config:\n"
        "  profile:\n"
        "    - channel:\n"
        "        - erasure_species: block-4-2\n"
        "          pdisk_category: 1\n"
        "          storage_pool_kind: ssd\n"
        "        - erasure_species: block-4-2\n"
        "          pdisk_category: 1\n"
        "          storage_pool_kind: ssd\n"
        "        - erasure_species: block-4-2\n"
        "          pdisk_category: 1\n"
        "          storage_pool_kind: ssd\n"
        "      profile_id: 0\n"
        "interconnect_config:\n"
        "  start_tcp: true\n"
        "  encryption_mode: OPTIONAL\n"
        "  path_to_certificate_file: \"/opt/ydb/certs/node.crt\"\n"
        "  path_to_private_key_file: \"/opt/ydb/certs/node.key\"\n"
        "  path_to_ca_file: \"/opt/ydb/certs/ca.crt\"\n"
        "grpc_config:\n"
        "  cert: \"/opt/ydb/certs/node.crt\"\n"
        "  key: \"/opt/ydb/certs/node.key\"\n"
        "  ca: \"/opt/ydb/certs/ca.crt\"\n"
        "  services_enabled:\n"
        "    - legacy\n";

        Y_ENSURE(Valid(v.Validate(yaml)));
    }

    Y_UNIT_TEST(MIRROR_3_DC_NODES) {
        auto v = StaticConfigBuilder().CreateValidator();
        const char* yaml = 
        "static_erasure: mirror-3-dc\n"
        "host_configs:\n"
        "- drive:\n"
        "  - path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "    type: SSD\n"
        "  - path: /dev/disk/by-partlabel/ydb_disk_ssd_02\n"
        "    type: SSD\n"
        "  - path: /dev/disk/by-partlabel/ydb_disk_ssd_03\n"
        "    type: SSD\n"
        "  host_config_id: 1\n"
        "hosts:\n"
        "- host: ydb-node-zone-a.local\n"
        "  host_config_id: 1                 # numeric host configuration template identifier\n"
        "  walle_location:                   # this parameter describes where host is located.\n"
        "    body: 1                         # string representing a host serial number.\n"
        "    data_center: 'zone-a'           # string representing the datacenter / availability zone where the host is located.\n"
        "                                    # if cluster is deployed using mirror-3-dc fault tolerance mode, all hosts must be distributed\n"
        "                                    # across 3 datacenters.\n"
        "    rack: '1'                       # string representing a rack identifier where the host is located.\n"
        "                                    # if cluster is deployed using block-4-2 erasure, all hosts should be distrubited\n"
        "                                    # accross at least 8 racks.\n"
        "- host: ydb-node-zone-b.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 2\n"
        "    data_center: 'zone-b'\n"
        "    rack: '2'\n"
        "- host: ydb-node-zone-c.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 3\n"
        "    data_center: 'zone-c'\n"
        "    rack: '3'\n"
        "domains_config:\n"
        "  # There can be only one root domain in a cluster. Domain name prefixes all scheme objects names, e.g. full name of a table table1 in database db1.\n"
        "  # in a cluster with domains_config.domain.name parameter set to Root would be equal to /Root/db1/table1\n"
        "  domain:\n"
        "  - name: Root\n"
        "    storage_pool_types:\n"
        "    - kind: ssd\n"
        "      pool_config:\n"
        "        box_id: 1\n"
        "        # fault tolerance mode name - none, block-4-2, or mirror-3-dc..\n"
        "        # See docs for more details https://ydb.tech/en/docs/deploy/configuration/config#domains-blob\n"
        "        erasure_species: mirror-3-dc\n"
        "        kind: ssd\n"
        "        geometry:\n"
        "          realm_level_begin: 10\n"
        "          realm_level_end: 20\n"
        "          domain_level_begin: 10\n"
        "          domain_level_end: 256\n"
        "        pdisk_filter:\n"
        "        - property:\n"
        "          - type: SSD  # device type to match host_configs.drive.type\n"
        "        vdisk_kind: Default\n"
        "  state_storage:\n"
        "  - ring:\n"
        "      node: [1, 2, 3]\n"
        "      nto_select: 3\n"
        "    ssid: 1\n"
        "table_service_config:\n"
        "  sql_version: 1\n"
        "actor_system_config:          # the configuration of the actor system which descibes how cores of the instance are distributed\n"
        "  executor:                   # accross different types of workloads in the instance.\n"
        "  - name: System              # system executor of the actor system. in this executor YDB launches system type of workloads, like system tablets\n"
        "                              # and reads from storage.\n"
        "    threads: 2                # the number of threads allocated to system executor.\n"
        "    type: BASIC\n"
        "  - name: User                # user executor of the actor system. In this executor YDB launches user workloads, like datashard activities,\n"
        "                              # queries and rpc calls.\n"
        "    threads: 3                # the number of threads allocated to user executor.\n"
        "    type: BASIC\n"
        "  - name: Batch               # user executor of the actor system. In this executor YDB launches batch operations, like scan queries, table\n"
        "                              # compactions, background compactions.\n"
        "    threads: 2                # the number of threads allocated to the batch executor.\n"
        "    type: BASIC\n"
        "  - name: IO                  # the io executor. In this executor launches sync operations and writes logs.\n"
        "    threads: 1\n"
        "    time_per_mailbox_micro_secs: 100\n"
        "    type: IO\n"
        "  - name: IC                  # the interconnect executor which YDB uses for network communications accross different nodes of the cluster.\n"
        "    spin_threshold: 10\n"
        "    threads: 1                # the number of threads allocated to the interconnect executor.\n"
        "    time_per_mailbox_micro_secs: 100\n"
        "    type: BASIC\n"
        "  scheduler:\n"
        "    progress_threshold: 10000\n"
        "    resolution: 256\n"
        "    spin_threshold: 0\n"
        "blob_storage_config:         # configuration of static blobstorage group.\n"
        "                             # YDB uses this group to store system tablets' data, like SchemeShard\n"
        "  service_set:\n"
        "    groups:\n"
        "    - erasure_species: mirror-3-dc # fault tolerance mode name for the static group\n"
        "      rings:          # in mirror-3-dc must have exactly 3 rings or availability zones\n"
        "      - fail_domains:  # first record: fail domains of the static group describe where each vdisk of the static group should be located.\n"
        "        - vdisk_locations:\n"
        "          - node_id: ydb-node-zone-a.local\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "        - vdisk_locations:\n"
        "          - node_id: ydb-node-zone-a.local\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_02\n"
        "        - vdisk_locations:\n"
        "          - node_id: ydb-node-zone-a.local\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_03\n"
        "      - fail_domains: # second ring: fail domains of the static group describe where each vdisk of the static group should be located.\n"
        "        - vdisk_locations:\n"
        "          - node_id: ydb-node-zone-b.local\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "        - vdisk_locations:\n"
        "          - node_id: ydb-node-zone-b.local\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_02\n"
        "        - vdisk_locations:\n"
        "          - node_id: ydb-node-zone-b.local\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_03\n"
        "      - fail_domains: # third ring: fail domains of the static group describe where each vdisk of the static group should be located.\n"
        "        - vdisk_locations:\n"
        "          - node_id: ydb-node-zone-c.local\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "        - vdisk_locations:\n"
        "          - node_id: ydb-node-zone-c.local\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_02\n"
        "        - vdisk_locations:\n"
        "          - node_id: ydb-node-zone-c.local\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_03\n"
        "channel_profile_config:\n"
        "  profile:\n"
        "  - channel:\n"
        "    - erasure_species: mirror-3-dc\n"
        "      pdisk_category: 0\n"
        "      storage_pool_kind: ssd\n"
        "    - erasure_species: mirror-3-dc\n"
        "      pdisk_category: 0\n"
        "      storage_pool_kind: ssd\n"
        "    - erasure_species: mirror-3-dc\n"
        "      pdisk_category: 0\n"
        "      storage_pool_kind: ssd\n"
        "    profile_id: 0\n"
        "interconnect_config:\n"
        "    start_tcp: true\n"
        "    encryption_mode: OPTIONAL\n"
        "    path_to_certificate_file: \"/opt/ydb/certs/node.crt\"\n"
        "    path_to_private_key_file: \"/opt/ydb/certs/node.key\"\n"
        "    path_to_ca_file: \"/opt/ydb/certs/ca.crt\"\n"
        "grpc_config:\n"
        "    cert: \"/opt/ydb/certs/node.crt\"\n"
        "    key: \"/opt/ydb/certs/node.key\"\n"
        "    ca: \"/opt/ydb/certs/ca.crt\"\n"
        "    services_enabled:\n"
        "    - legacy\n";

        Y_ENSURE(Valid(v.Validate(yaml)));
    }

    Y_UNIT_TEST(MIRROR_3_DC_NODES_IN_MEMORY) {
        auto v = StaticConfigBuilder().CreateValidator();
        const char* yaml = 
        "static_erasure: mirror-3-dc\n"
        "host_configs:\n"
        "- drive:\n"
        "  - path: SectorMap:1:64\n"
        "    type: SSD\n"
        "  - path: SectorMap:2:64\n"
        "    type: SSD\n"
        "  - path: SectorMap:3:64\n"
        "    type: SSD\n"
        "  host_config_id: 1\n"
        "hosts:\n"
        "- host: localhost\n"
        "  host_config_id: 1\n"
        "  port: 19001\n"
        "  walle_location:\n"
        "    body: 1\n"
        "    data_center: '1'\n"
        "    rack: '1'\n"
        "- host: localhost\n"
        "  host_config_id: 1\n"
        "  port: 19002\n"
        "  walle_location:\n"
        "    body: 2\n"
        "    data_center: '2'\n"
        "    rack: '2'\n"
        "- host: localhost\n"
        "  host_config_id: 1\n"
        "  port: 19003\n"
        "  walle_location:\n"
        "    body: 3\n"
        "    data_center: '3'\n"
        "    rack: '3'\n"
        "domains_config:\n"
        "  domain:\n"
        "  - name: Root\n"
        "    storage_pool_types:\n"
        "    - kind: ssd\n"
        "      pool_config:\n"
        "        box_id: 1\n"
        "        erasure_species: mirror-3-dc\n"
        "        kind: ssd\n"
        "        geometry:\n"
        "          realm_level_begin: 10\n"
        "          realm_level_end: 20\n"
        "          domain_level_begin: 10\n"
        "          domain_level_end: 256\n"
        "        pdisk_filter:\n"
        "        - property:\n"
        "          - type: SSD\n"
        "        vdisk_kind: Default\n"
        "  state_storage:\n"
        "  - ring:\n"
        "      node: [1, 2, 3]\n"
        "      nto_select: 3\n"
        "    ssid: 1\n"
        "table_service_config:\n"
        "  sql_version: 1\n"
        "actor_system_config:\n"
        "  executor:\n"
        "  - name: System\n"
        "    spin_threshold: 0\n"
        "    threads: 2\n"
        "    type: BASIC\n"
        "  - name: User\n"
        "    spin_threshold: 0\n"
        "    threads: 3\n"
        "    type: BASIC\n"
        "  - name: Batch\n"
        "    spin_threshold: 0\n"
        "    threads: 2\n"
        "    type: BASIC\n"
        "  - name: IO\n"
        "    threads: 1\n"
        "    time_per_mailbox_micro_secs: 100\n"
        "    type: IO\n"
        "  - name: IC\n"
        "    spin_threshold: 10\n"
        "    threads: 1\n"
        "    time_per_mailbox_micro_secs: 100\n"
        "    type: BASIC\n"
        "  scheduler:\n"
        "    progress_threshold: 10000\n"
        "    resolution: 256\n"
        "    spin_threshold: 0\n"
        "blob_storage_config:\n"
        "  service_set:\n"
        "    groups:\n"
        "    - erasure_species: mirror-3-dc\n"
        "      rings:\n"
        "      - fail_domains:\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"localhost:19001\"\n"
        "            pdisk_category: SSD\n"
        "            path: SectorMap:1:64\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"localhost:19001\"\n"
        "            pdisk_category: SSD\n"
        "            path: SectorMap:2:64\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"localhost:19001\"\n"
        "            pdisk_category: SSD\n"
        "            path: SectorMap:3:64\n"
        "      - fail_domains:\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"localhost:19002\"\n"
        "            pdisk_category: SSD\n"
        "            path: SectorMap:1:64\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"localhost:19002\"\n"
        "            pdisk_category: SSD\n"
        "            path: SectorMap:2:64\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"localhost:19002\"\n"
        "            pdisk_category: SSD\n"
        "            path: SectorMap:3:64\n"
        "      - fail_domains:\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"localhost:19003\"\n"
        "            pdisk_category: SSD\n"
        "            path: SectorMap:1:64\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"localhost:19003\"\n"
        "            pdisk_category: SSD\n"
        "            path: SectorMap:2:64\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"localhost:19003\"\n"
        "            pdisk_category: SSD\n"
        "            path: SectorMap:3:64\n"
        "channel_profile_config:\n"
        "  profile:\n"
        "  - channel:\n"
        "    - erasure_species: mirror-3-dc\n"
        "      pdisk_category: 1\n"
        "      storage_pool_kind: ssd\n"
        "    - erasure_species: mirror-3-dc\n"
        "      pdisk_category: 1\n"
        "      storage_pool_kind: ssd\n"
        "    - erasure_species: mirror-3-dc\n"
        "      pdisk_category: 1\n"
        "      storage_pool_kind: ssd\n"
        "    profile_id: 0\n";
        
        Y_ENSURE(Valid(v.Validate(yaml)));
    }

    Y_UNIT_TEST(MIRROR_3_DC_9_NODES) {
        auto v = StaticConfigBuilder().CreateValidator();
        const char* yaml = 
        "static_erasure: mirror-3-dc\n"
        "host_configs:\n"
        "- drive:\n"
        "  - path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "    type: SSD\n"
        "  - path: /dev/disk/by-partlabel/ydb_disk_ssd_02\n"
        "    type: SSD\n"
        "  host_config_id: 1\n"
        "hosts:\n"
        "- host: ydb-node-zone-a-1.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 1\n"
        "    data_center: 'zone-a'\n"
        "    rack: '1'\n"
        "- host: ydb-node-zone-a-2.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 2\n"
        "    data_center: 'zone-a'\n"
        "    rack: '2'\n"
        "- host: ydb-node-zone-a-3.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 3\n"
        "    data_center: 'zone-a'\n"
        "    rack: '3'\n"
        "\n"
        "- host: ydb-node-zone-b-1.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 4\n"
        "    data_center: 'zone-b'\n"
        "    rack: '4'\n"
        "- host: ydb-node-zone-b-2.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 5\n"
        "    data_center: 'zone-b'\n"
        "    rack: '5'\n"
        "- host: ydb-node-zone-b-3.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 6\n"
        "    data_center: 'zone-b'\n"
        "    rack: '6'\n"
        "\n"
        "- host: ydb-node-zone-c-1.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 7\n"
        "    data_center: 'zone-c'\n"
        "    rack: '7'\n"
        "- host: ydb-node-zone-c-2.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 8\n"
        "    data_center: 'zone-c'\n"
        "    rack: '8'\n"
        "- host: ydb-node-zone-c-3.local\n"
        "  host_config_id: 1\n"
        "  walle_location:\n"
        "    body: 9\n"
        "    data_center: 'zone-c'\n"
        "    rack: '9'\n"
        "\n"
        "domains_config:\n"
        "  domain:\n"
        "  - name: Root\n"
        "    storage_pool_types:\n"
        "    - kind: ssd\n"
        "      pool_config:\n"
        "        box_id: 1\n"
        "        erasure_species: mirror-3-dc\n"
        "        kind: ssd\n"
        "        pdisk_filter:\n"
        "        - property:\n"
        "          - type: SSD\n"
        "        vdisk_kind: Default\n"
        "  state_storage:\n"
        "  - ring:\n"
        "      node: [1, 2, 3, 4, 5, 6, 7, 8, 9]\n"
        "      nto_select: 9\n"
        "    ssid: 1\n"
        "table_service_config:\n"
        "  sql_version: 1\n"
        "actor_system_config:\n"
        "  executor:\n"
        "  - name: System\n"
        "    spin_threshold: 0\n"
        "    threads: 2\n"
        "    type: BASIC\n"
        "  - name: User\n"
        "    spin_threshold: 0\n"
        "    threads: 3\n"
        "    type: BASIC\n"
        "  - name: Batch\n"
        "    spin_threshold: 0\n"
        "    threads: 2\n"
        "    type: BASIC\n"
        "  - name: IO\n"
        "    threads: 1\n"
        "    time_per_mailbox_micro_secs: 100\n"
        "    type: IO\n"
        "  - name: IC\n"
        "    spin_threshold: 10\n"
        "    threads: 1\n"
        "    time_per_mailbox_micro_secs: 100\n"
        "    type: BASIC\n"
        "  scheduler:\n"
        "    progress_threshold: 10000\n"
        "    resolution: 256\n"
        "    spin_threshold: 0\n"
        "blob_storage_config:\n"
        "  service_set:\n"
        "    groups:\n"
        "    - erasure_species: mirror-3-dc\n"
        "      rings:\n"
        "      - fail_domains:\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"ydb-node-zone-a-1.local\"\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"ydb-node-zone-a-2.local\"\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"ydb-node-zone-a-3.local\"\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "      - fail_domains:\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"ydb-node-zone-b-1.local\"\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"ydb-node-zone-b-2.local\"\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"ydb-node-zone-b-3.local\"\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "      - fail_domains:\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"ydb-node-zone-c-1.local\"\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"ydb-node-zone-c-2.local\"\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "        - vdisk_locations:\n"
        "          - node_id: \"ydb-node-zone-c-3.local\"\n"
        "            pdisk_category: SSD\n"
        "            path: /dev/disk/by-partlabel/ydb_disk_ssd_01\n"
        "channel_profile_config:\n"
        "  profile:\n"
        "  - channel:\n"
        "    - erasure_species: mirror-3-dc\n"
        "      pdisk_category: 1\n"
        "      storage_pool_kind: ssd\n"
        "    - erasure_species: mirror-3-dc\n"
        "      pdisk_category: 1\n"
        "      storage_pool_kind: ssd\n"
        "    - erasure_species: mirror-3-dc\n"
        "      pdisk_category: 1\n"
        "      storage_pool_kind: ssd\n"
        "    profile_id: 0\n"
        "interconnect_config:\n"
        "    start_tcp: true\n"
        "    encryption_mode: OPTIONAL\n"
        "    path_to_certificate_file: \"/opt/ydb/certs/node.crt\"\n"
        "    path_to_private_key_file: \"/opt/ydb/certs/node.key\"\n"
        "    path_to_ca_file: \"/opt/ydb/certs/ca.crt\"\n"
        "grpc_config:\n"
        "    cert: \"/opt/ydb/certs/node.crt\"\n"
        "    key: \"/opt/ydb/certs/node.key\"\n"
        "    ca: \"/opt/ydb/certs/ca.crt\"\n"
        "    services_enabled:\n"
        "    - legacy\n";
        
        Y_ENSURE(Valid(v.Validate(yaml)));
    }

    Y_UNIT_TEST(SINGLE_NODE_IN_MEMORY) {
        auto v = StaticConfigBuilder().CreateValidator();
        const char* yaml = 
        "static_erasure: none\n"
        "host_configs:\n"
        "- drive:\n"
        "  - path: SectorMap:1:64\n"
        "    type: SSD\n"
        "  host_config_id: 1\n"
        "hosts:\n"
        "- host: localhost\n"
        "  host_config_id: 1\n"
        "  port: 19001\n"
        "  walle_location:\n"
        "    body: 1\n"
        "    data_center: '1'\n"
        "    rack: '1'\n"
        "domains_config:\n"
        "  domain:\n"
        "  - name: Root\n"
        "    storage_pool_types:\n"
        "    - kind: ssd\n"
        "      pool_config:\n"
        "        box_id: 1\n"
        "        erasure_species: none\n"
        "        kind: ssd\n"
        "        pdisk_filter:\n"
        "        - property:\n"
        "          - type: SSD\n"
        "        vdisk_kind: Default\n"
        "  state_storage:\n"
        "  - ring:\n"
        "      node:\n"
        "      - 1\n"
        "      nto_select: 1\n"
        "    ssid: 1\n"
        "table_service_config:\n"
        "  sql_version: 1\n"
        "actor_system_config:\n"
        "  executor:\n"
        "  - name: System\n"
        "    spin_threshold: 0\n"
        "    threads: 2\n"
        "    type: BASIC\n"
        "  - name: User\n"
        "    spin_threshold: 0\n"
        "    threads: 3\n"
        "    type: BASIC\n"
        "  - name: Batch\n"
        "    spin_threshold: 0\n"
        "    threads: 2\n"
        "    type: BASIC\n"
        "  - name: IO\n"
        "    threads: 1\n"
        "    time_per_mailbox_micro_secs: 100\n"
        "    type: IO\n"
        "  - name: IC\n"
        "    spin_threshold: 10\n"
        "    threads: 1\n"
        "    time_per_mailbox_micro_secs: 100\n"
        "    type: BASIC\n"
        "  scheduler:\n"
        "    progress_threshold: 10000\n"
        "    resolution: 256\n"
        "    spin_threshold: 0\n"
        "blob_storage_config:\n"
        "  service_set:\n"
        "    groups:\n"
        "    - erasure_species: none\n"
        "      rings:\n"
        "      - fail_domains:\n"
        "        - vdisk_locations:\n"
        "          - node_id: 1\n"
        "            path: SectorMap:1:64\n"
        "            pdisk_category: SSD\n"
        "channel_profile_config:\n"
        "  profile:\n"
        "  - channel:\n"
        "    - erasure_species: none\n"
        "      pdisk_category: 1\n"
        "      storage_pool_kind: ssd\n"
        "    - erasure_species: none\n"
        "      pdisk_category: 1\n"
        "      storage_pool_kind: ssd\n"
        "    - erasure_species: none\n"
        "      pdisk_category: 1\n"
        "      storage_pool_kind: ssd\n"
        "    profile_id: 0\n";
        
        Y_ENSURE(Valid(v.Validate(yaml)));
    }

}
