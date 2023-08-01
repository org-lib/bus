package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"strings"
)

type Info struct {
	Host     string
	Port     int
	DB       int
	Password string
}

func NewClient(cnf Info) *redis.Client {
	if cnf.Host == "" {
		fmt.Println("使用了默认地址：localhost")
		cnf.Host = "localhost"
	}
	if cnf.Port < 0 {
		fmt.Println("使用了默认端口：6379")
		cnf.Port = 6379
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%d", cnf.Host, cnf.Port),
		Password: cnf.Password,
		DB:       cnf.DB,
	})
	return client
}

/*
通过key["xx"]["yy"]直接获取，或者通过二级key[xx]获取一组对象
info: # Server
redis_version:4.0.11
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:4b1fcccca6e8040b
redis_mode:standalone
os:Linux 2.6.32-642.el6.x86_64 x86_64
arch_bits:64
multiplexing_api:epoll
atomicvar_api:sync-builtin
gcc_version:4.4.7
process_id:21895
run_id:ee4c85f05da5aa09e5916eeb571d8822d5d9ed35
tcp_port:27023
uptime_in_seconds:86822734
uptime_in_days:1004
hz:10
lru_clock:1341849
executable:/home/xxx/redis-server
config_file:/tmp.conf

# Clients
connected_clients:36
client_longest_output_list:0
client_biggest_input_buf:0
blocked_clients:17

# Memory
used_memory:12432808
used_memory_human:11.86M
used_memory_rss:22921216
used_memory_rss_human:21.86M
used_memory_peak:25306112
used_memory_peak_human:24.13M
used_memory_peak_perc:49.13%
used_memory_overhead:11995374
used_memory_startup:786528
used_memory_dataset:437434
used_memory_dataset_perc:3.76%
total_system_memory:8254570496
total_system_memory_human:7.69G
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:268435456
maxmemory_human:256.00M
maxmemory_policy:volatile-lru
mem_fragmentation_ratio:1.84
mem_allocator:jemalloc-4.0.3
active_defrag_running:0
lazyfree_pending_objects:0

# Persistence
loading:0
rdb_changes_since_last_save:3182887
rdb_bgsave_in_progress:0
rdb_last_save_time:1640670601
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_last_cow_size:4714496
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_last_write_status:ok
aof_last_cow_size:4661248
aof_current_size:45472028
aof_base_size:6617
aof_pending_rewrite:0
aof_buffer_length:0
aof_rewrite_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:579

# Stats
total_connections_received:6339915
total_commands_processed:727862031
instantaneous_ops_per_sec:12
total_net_input_bytes:55293516593
total_net_output_bytes:250841368254
instantaneous_input_kbps:0.49
instantaneous_output_kbps:4.72
rejected_connections:0
sync_full:1
sync_partial_ok:0
sync_partial_err:1
expired_keys:74
expired_stale_perc:0.00
expired_time_cap_reached_count:0
evicted_keys:0
keyspace_hits:1584973
keyspace_misses:138809
pubsub_channels:3
pubsub_patterns:2
latest_fork_usec:1290
migrate_cached_sockets:0
slave_expires_tracked_keys:0
active_defrag_hits:0
active_defrag_misses:0
active_defrag_key_hits:0
active_defrag_key_misses:0

# Replication
role:master
connected_slaves:1
slave0:ip=10.88.27.177,port=27021,state=online,offset=30528115729,lag=0
master_replid:91075bf672b2442121934fc5d870fd6052d04427
master_replid2:f8a3e915b77a895d6af5de9d1df4a63e49dba3cc
master_repl_offset:30528115729
second_repl_offset:29135280300
repl_backlog_active:1
repl_backlog_size:10000000
repl_backlog_first_byte_offset:30518115730
repl_backlog_histlen:10000000

# CPU
used_cpu_sys:61603.61
used_cpu_user:39983.96
used_cpu_sys_children:0.35
used_cpu_user_children:0.18

# Cluster
cluster_enabled:0

# Keyspace
db2:keys=12,expires=0,avg_ttl=0
*/
func InfoToMap(client *redis.Client) (map[string]map[string]interface{}, error) {
	msg, err := client.Info().Result()
	if err != nil {
		return nil, err
	}
	info := make(map[string]map[string]interface{})
	subInfo := make(map[string]interface{})

	// one end

	info_title := ""
	for _, s := range strings.Split(string(msg), "\r\n") {
		if strings.Trim(s, " ") == "" {
			continue
		}
		if strings.HasPrefix(s, "# ") {
			if info_title != "" {
				info[info_title] = subInfo

				//重置子 map

				subInfo = nil
				subInfo = make(map[string]interface{})
			}
			info_title = strings.ReplaceAll(s, "# ", "")
			continue
		}
		kv := strings.Split(s, ":")
		subInfo[kv[0]] = kv[1]
	}
	// 处理最后一个map，除非Redis info 命令结果集，最后一行不是空行或者nil的标识

	info[info_title] = subInfo

	//重置子 map

	subInfo = nil

	info_title = ""

	return info, nil
}

/*
通过key["xx"]直接获取
info: # Server
redis_version:4.0.11
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:4b1fcccca6e8040b
redis_mode:standalone
os:Linux 2.6.32-642.el6.x86_64 x86_64
arch_bits:64
multiplexing_api:epoll
atomicvar_api:sync-builtin
gcc_version:4.4.7
process_id:21895
run_id:ee4c85f05da5aa09e5916eeb571d8822d5d9ed35
tcp_port:27023
uptime_in_seconds:86822734
uptime_in_days:1004
hz:10
lru_clock:1341849
executable:/home/xxx/redis-server
config_file:/tmp.conf

# Clients
connected_clients:36
client_longest_output_list:0
client_biggest_input_buf:0
blocked_clients:17

# Memory
used_memory:12432808
used_memory_human:11.86M
used_memory_rss:22921216
used_memory_rss_human:21.86M
used_memory_peak:25306112
used_memory_peak_human:24.13M
used_memory_peak_perc:49.13%
used_memory_overhead:11995374
used_memory_startup:786528
used_memory_dataset:437434
used_memory_dataset_perc:3.76%
total_system_memory:8254570496
total_system_memory_human:7.69G
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:268435456
maxmemory_human:256.00M
maxmemory_policy:volatile-lru
mem_fragmentation_ratio:1.84
mem_allocator:jemalloc-4.0.3
active_defrag_running:0
lazyfree_pending_objects:0

# Persistence
loading:0
rdb_changes_since_last_save:3182887
rdb_bgsave_in_progress:0
rdb_last_save_time:1640670601
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_last_cow_size:4714496
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_last_write_status:ok
aof_last_cow_size:4661248
aof_current_size:45472028
aof_base_size:6617
aof_pending_rewrite:0
aof_buffer_length:0
aof_rewrite_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:579

# Stats
total_connections_received:6339915
total_commands_processed:727862031
instantaneous_ops_per_sec:12
total_net_input_bytes:55293516593
total_net_output_bytes:250841368254
instantaneous_input_kbps:0.49
instantaneous_output_kbps:4.72
rejected_connections:0
sync_full:1
sync_partial_ok:0
sync_partial_err:1
expired_keys:74
expired_stale_perc:0.00
expired_time_cap_reached_count:0
evicted_keys:0
keyspace_hits:1584973
keyspace_misses:138809
pubsub_channels:3
pubsub_patterns:2
latest_fork_usec:1290
migrate_cached_sockets:0
slave_expires_tracked_keys:0
active_defrag_hits:0
active_defrag_misses:0
active_defrag_key_hits:0
active_defrag_key_misses:0

# Replication
role:master
connected_slaves:1
slave0:ip=10.88.27.177,port=27021,state=online,offset=30528115729,lag=0
master_replid:91075bf672b2442121934fc5d870fd6052d04427
master_replid2:f8a3e915b77a895d6af5de9d1df4a63e49dba3cc
master_repl_offset:30528115729
second_repl_offset:29135280300
repl_backlog_active:1
repl_backlog_size:10000000
repl_backlog_first_byte_offset:30518115730
repl_backlog_histlen:10000000

# CPU
used_cpu_sys:61603.61
used_cpu_user:39983.96
used_cpu_sys_children:0.35
used_cpu_user_children:0.18

# Cluster
cluster_enabled:0

# Keyspace
db2:keys=12,expires=0,avg_ttl=0
*/
func Info2Map(client *redis.Client) (map[string]interface{}, error) {
	msg, err := client.Info().Result()
	if err != nil {
		return nil, err
	}
	info := make(map[string]interface{})

	// one end

	for _, s := range strings.Split(string(msg), "\r\n") {
		if strings.Trim(s, " ") == "" || strings.HasPrefix(s, "# ") {
			continue
		}
		kv := strings.Split(s, ":")
		info[kv[0]] = kv[1]
	}
	return info, nil
}
