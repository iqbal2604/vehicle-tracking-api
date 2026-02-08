# TODO: Add User Roles (Admin and Driver) to Vehicle Tracking API

## Steps Completed

1. **Update User Model** ✅:
   - Added `Role` field to `models/user.go` (string type, default "driver" via GORM tag).

2. **Update User DTO** ✅:
   - Modified `dtos/user_dto.go` to include the `Role` field in the `UserResponse` struct and `ToUserResponse` function.

3. **Update Register Request** ✅:
   - Added `Role` field to `requests/register_request.go` (optional, defaults to "driver").

4. **Update User Service** ✅:
   - Modified `services/user_service.go` Register method to accept a `role` parameter and set it in the user model.

5. **Update Auth Handler** ✅:
   - Updated `handlers/auth_handler.go` Register method to parse the `Role` from request body and pass it to the service.

6. **Update Vehicle Handler** ✅:
   - Added role-based access checks in `handlers/vehicle_handler.go` for all CRUD operations:
     - `CreateVehicle`: Drivers set UserID to their own; admins can set any.
     - `GetVehicle`: Drivers can only access their own vehicles; admins can access all.
     - `ListVehicle`: Drivers list their own; admins list all vehicles.
     - `UpdateVehicle`: Added ownership check before updating.
     - `DeleteVehicle`: Added ownership check before deleting.
   - Added `getUserRole` helper and updated constructor to include `UserRepository`.

7. **Update Vehicle Service** ✅:
   - Added `ListAllVehicles` method to list all vehicles for admins.

8. **Update main.go** ✅:
   - Modified `VehicleHandler` instantiation to pass `userRepo`.

9. **Database Migration** ✅:
   - Ran `go run cmd/main.go` successfully; GORM auto-migrated the `role` column to the `users` table.

## Explanation of Changes
- **Role Logic**: "driver" users can only manage their own vehicles. "admin" users can manage all vehicles.
- **Registration**: Role is optional in request; defaults to "driver". Admins can be created by including `"role": "admin"` in the register payload.
- **Access Control**: Each vehicle operation now checks the user's role and enforces ownership for drivers.
- **Database**: The `role` column was added with a default value of "driver".
- **No Breaking Changes**: Existing functionality for drivers remains the same; admins have expanded access.

## Followup Steps
- Run `go mod tidy` (no new dependencies added).
- Test endpoints using `test.http`:
  - Register a driver: `POST /api/register` with `{"name":"Driver","email":"driver@example.com","password":"pass"}`
  - Register an admin: `POST /api/register` with `{"name":"Admin","email":"admin@example.com","password":"pass","role":"admin"}`
  - Login and test vehicle CRUD with different roles to verify access restrictions.
