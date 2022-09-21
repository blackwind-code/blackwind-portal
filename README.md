# blackwind-portal
Blackwind Service Portal for account registration &amp; VPN device management  
DGIST students can make blackwind service account,
and connect their devices to our VPN system through this portal  
  
# Reference
[Pocket Base Doc](https://pocketbase.io/docs/)

# Install & Setup
## Build
```bash
git clone https://github.com/blackwind-code/blackwind-portal.git
cd blackwind-portal/cmd/blackwind-portal
go build main.go
```  
- You need golang installed on your system  
- You need [blackwind-portal-driver] up and running

# Run
```bash
export SECRET=<secret-password>

export VPN_DRIVER_URL=<blackwind-portal-driver-ip:port>
export OPENSTACK_DRIVER_URL=<blackwind-portal-driver-ip:port>

cd blackwind-portal/cmd/blackwind-portal/
./main serve
```

# Set Database Schema
Register admin account on first startup (http://127.0.0.1:8090/_/)  
Admin page: Settings(on the left menu bar) -> Import collections -> Load from JSON file -> import "config/pb_schema.json"  
