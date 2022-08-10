# HW 3 REST API client implementation
### https://gorest.co.in/  *GraphQL and REST API for Testing and Prototyping*

example\main.go
```go
	client := api.NewUserApiClient(shared.NewConfig()) //use default config with default token
	ctx := context.Background()

	println("GET ALL ")
	users, err := client.GetAll(ctx)
	if err != nil {
		println("ERROR ", err.Error())
	}

	prettyPrint(users)

	user := users[0]
	trace := user

	println("GET BY ID : ", user.Id)
	if trace, err = client.GetById(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}

	prettyPrint(trace)

	println("UPDATE NAME\\EMAIL for ID : ", user.Id, " name : ", user.Name)
	if err := client.Update(ctx, user.Id,
		domain.User{
			Name:  user.Name + time.Now().Format(time.RFC3339),
			Email: "NEW-EMAIL-" + user.Email,
		}); err != nil {
		println("ERROR ", err.Error())
	}

	if trace, err = client.GetById(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}
	prettyPrint(trace)

	println("UPDATE STATUS ID : ", user.Id, " status : ", user.Status, "  email : ", user.Email)
	if err := client.UpdateStatus(ctx, user.Id, getNextStatus(user.Status)); err != nil {
		println("ERROR ", err.Error())
	}

	if trace, err = client.GetById(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}
	prettyPrint(trace)

	println("UPDATE GENDER ID : ", user.Id, " gender : ", user.Gender)
	if err := client.UpdateGender(ctx, user.Id, getNextGender(user.Gender)); err != nil {
		println("ERROR ", err.Error())
	}

	if trace, err = client.GetById(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}
	prettyPrint(trace)

	println("DELETE USER ID : ", user.Id)
	if err := client.Delete(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}

	if trace, err = client.GetById(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}
	prettyPrint(trace)
```
```text
GET ALL
time="2022-08-10T18:32:02+03:00" level=debug msg="requesting... GET https://gorest.co.in/public/v2/users"
time="2022-08-10T18:32:02+03:00" level=debug msg="200 OK -1"
[
        {
                "id": 3389,
                "name": "Mr. Gopee Somayaji",
                "email": "mr_gopee_somayaji@wisoky-prosacco.biz",
                "gender": "male",
                "status": "active"
        },
        {
                "id": 3388,
                "name": "Nanda Gowda",
                "email": "gowda_nanda@corkery.biz",
                "gender": "male",
                "status": "active"
        },
        {
                "id": 3387,
                "name": "Giriraj Chopra",
                "email": "giriraj_chopra@lesch-kassulke.net",
                "gender": "male",
                "status": "active"
        },
        {
                "id": 3386,
                "name": "Msgr. Mohini Jain",
                "email": "msgr_jain_mohini@orn.co",
                "gender": "female",
                "status": "active"
        },
        {
                "id": 3385,
                "name": "Chandira Sinha",
                "email": "chandira_sinha@ward.com",
                "gender": "male",
                "status": "inactive"
        },
        {
                "id": 3384,
                "name": "Tanushree Gupta",
                "email": "gupta_tanushree@mcdermott.net",
                "gender": "female",
                "status": "inactive"
        },
        {
                "id": 3383,
                "name": "Damayanti Sharma",
                "email": "sharma_damayanti@schulist.com",
                "gender": "male",
                "status": "active"
        },
        {
                "id": 3382,
                "name": "The Hon. Divjot Rana",
                "email": "hon_divjot_the_rana@howe.com",
                "gender": "female",
                "status": "active"
        },
        {
                "id": 3381,
                "name": "Kartik Chaturvedi",
                "email": "kartik_chaturvedi@stiedemann.name",
                "gender": "male",
                "status": "active"
        },
        {
                "id": 3380,
                "name": "Fr. Vrund Patil",
                "email": "fr_patil_vrund@konopelski-franecki.info",
                "gender": "male",
                "status": "active"
        }
]
GET BY ID :  3389
time="2022-08-10T18:32:02+03:00" level=debug msg="requesting... GET https://gorest.co.in/public/v2/users/3389"
time="2022-08-10T18:32:03+03:00" level=debug msg="200 OK -1"
{
        "id": 3389,
        "name": "Mr. Gopee Somayaji",
        "email": "mr_gopee_somayaji@wisoky-prosacco.biz",
        "gender": "male",
        "status": "active"
}
UPDATE NAME\EMAIL for ID :  3389  name :  Mr. Gopee Somayaji
time="2022-08-10T18:32:03+03:00" level=debug msg="requesting... PUT https://gorest.co.in/public/v2/users/3389"
time="2022-08-10T18:32:03+03:00" level=debug msg="200 OK -1"
time="2022-08-10T18:32:03+03:00" level=debug msg="requesting... GET https://gorest.co.in/public/v2/users/3389"
time="2022-08-10T18:32:03+03:00" level=debug msg="200 OK -1"
{
        "id": 3389,
        "name": "Mr. Gopee Somayaji2022-08-10T18:32:03+03:00",
        "email": "NEW-EMAIL-mr_gopee_somayaji@wisoky-prosacco.biz",
        "gender": "male",
        "status": "active"
}
UPDATE STATUS ID :  3389  status :  active   email :  mr_gopee_somayaji@wisoky-prosacco.biz
time="2022-08-10T18:32:03+03:00" level=debug msg="requesting... PUT https://gorest.co.in/public/v2/users/3389"
time="2022-08-10T18:32:03+03:00" level=debug msg="200 OK -1"
time="2022-08-10T18:32:03+03:00" level=debug msg="requesting... GET https://gorest.co.in/public/v2/users/3389"
time="2022-08-10T18:32:03+03:00" level=debug msg="200 OK -1"
{
        "id": 3389,
        "name": "Mr. Gopee Somayaji2022-08-10T18:32:03+03:00",
        "email": "NEW-EMAIL-mr_gopee_somayaji@wisoky-prosacco.biz",
        "gender": "male",
        "status": "inactive"
}
UPDATE GENDER ID :  3389  gender :  male
time="2022-08-10T18:32:03+03:00" level=debug msg="requesting... PATCH https://gorest.co.in/public/v2/users/3389"
time="2022-08-10T18:32:04+03:00" level=debug msg="200 OK -1"
time="2022-08-10T18:32:04+03:00" level=debug msg="requesting... GET https://gorest.co.in/public/v2/users/3389"
time="2022-08-10T18:32:04+03:00" level=debug msg="200 OK -1"
{
        "id": 3389,
        "name": "Mr. Gopee Somayaji2022-08-10T18:32:03+03:00",
        "email": "NEW-EMAIL-mr_gopee_somayaji@wisoky-prosacco.biz",
        "gender": "female",
        "status": "inactive"
}
DELETE USER ID :  3389
time="2022-08-10T18:32:04+03:00" level=debug msg="requesting... DELETE https://gorest.co.in/public/v2/users/3389"
time="2022-08-10T18:32:04+03:00" level=debug msg="204 No Content 0"
time="2022-08-10T18:32:04+03:00" level=debug msg="requesting... GET https://gorest.co.in/public/v2/users/3389"
time="2022-08-10T18:32:04+03:00" level=debug msg="404 Not Found -1"
{}
```