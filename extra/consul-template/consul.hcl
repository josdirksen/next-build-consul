max_stale = "10m"
retry     = "10s"
wait      = "5s:20s"

template {
  source = "/etc/haproxy/haproxy.template"
  destination = "/etc/haproxy/haproxy.cfg"
  command = "/hap.sh"
  perms = 0600
}
