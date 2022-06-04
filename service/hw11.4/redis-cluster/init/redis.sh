sleep 10 && redis-cli --cluster create 111.111.0.11:7011 111.111.0.21:7021 111.111.0.31:7031 111.111.0.12:7012 111.111.0.22:7022 111.111.0.32:7032 --cluster-replicas 1 --cluster-yes
 