
1. **Requirement**
   * install `Docker` latest or stable version
   * Install `Make` package latest or stable version
   * Install `Postman`

2. **Install api on docker**
   1. CD to root folder `api-searching`
   2. Type `make install`
   3. Type `make start`
   4. Type `make logs`  
      Output Sucess  
      `... 2023/09/25 09:16:19 stdout: [GIN-debug] Listening and serving HTTP on :8011`

3. **Execute test**
   - Rest API: Using postman
      1. URL: `http://localhost:8080/v1/users/1`
      2. Header request: _data any(just make it as sample)_  
         `X-PLATFORM:IOS`  
         `X-DEVICE-TYPE:phone`  
         `X-DEVICE-ID:620005a5-1305-4668-9fb2-3ba250a57ab9`  
         `X-LANG:en`  
         `X-CHANNEL:2`

   - Unit test: Using `make` file
      1. CD to root folder `api-searching`.
      2. Test controller `make unit-test-controller` see report html at `tests/report/user.controller.coverage.html`
      3. Test Service `make unit-test-service`  see report html at `tests/report/user.service.coverage.html`
   
