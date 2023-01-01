io_mode = "async"

service "http" "web_proxy" {
  listen_addr = "127.0.0.1:8080"
  
  process "main" {
    command = csvdecode("command\targument\n/usr/local/bin/awesome-app\tserver", "\t")
  }

  process "mgmt" {
    command = csvdecode("command|argument\n/usr/local/bin/awesome-app|mgmt", "|")

  }
}