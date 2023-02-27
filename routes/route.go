package routes

import (
	"systemtest2/controllers"
	mddlwr "systemtest2/middleware"
	"systemtest2/salescontroller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With",
		ExposeHeaders:    "Origin",
	}))
	//------------>
	//------------------------->
	//---------------------------------->
	//--------------------------------------------->
	//------------------------Create Route Area------------------------------------>
	//------------------------10 Inside Location----------------------------------->
	app.Post("/api/focreate", controllers.FOCreate)        //1
	app.Post("/api/mrcreate", controllers.MRCreate)        //2
	app.Post("/api/mrtcreate", controllers.MRTICreate)     //3
	app.Post("/api/rscreate", controllers.RSCreate)        //4
	app.Post("/api/kkdcreate", controllers.KiskendaCreate) //5
	app.Post("/api/mlicreate", controllers.MaliawanCreate) //6
	app.Post("/api/occreate", controllers.OCafeCreate)     //7
	app.Post("/api/docreate", controllers.DOCreate)        //8
	app.Post("/api/vlacreate", controllers.VillaCreate)    //9
	app.Post("/api/krkcreate", controllers.KaraokeCreate)  //10
	//------------------------- 2 Outside Location -------------------------------->
	app.Post("/api/pbcreate", controllers.PBCreate)  //11
	app.Post("/api/toscreate", controllers.TOCreate) //12
	//------------------------End Create Route Area-------------------------------->
	//--------------------------------------------------->
	//------------------------User Cek Route Area---------------------------------->
	//------------------------10 Inside Location----------------------------------->
	app.Get("api/focek", controllers.FOUserCek)        //1
	app.Get("api/mrcek", controllers.MRUserCek)        //2
	app.Get("api/mrtcek", controllers.MRTIUserCek)     //3
	app.Get("api/rscek", controllers.RamaSintaUserCek) //4
	app.Get("api/kkdcek", controllers.KiskendaUserCek) //5
	app.Get("api/mlicek", controllers.MaliawanUserCek) //6
	app.Get("api/occek", controllers.OCafeUserCek)     //7
	app.Get("api/docek", controllers.DOUserCek)        //8
	app.Get("api/vlacek", controllers.VillaUserCek)    //9
	app.Get("api/krkcek", controllers.KaraokeUserCek)  //10
	//------------------------- 2 Outside Location -------------------------------->
	app.Get("/api/pbcek", controllers.PBUserCek)   //11
	app.Get("/api/toscek", controllers.TOSUserCek) //12
	//------------------------End User Cek Route Area------------------------------>

	//--------------------------------------------------->
	//------------------------All User Route Area---------------------------------->
	//------------------------10 Inside Location----------------------------------->
	app.Get("api/foall", mddlwr.IsAuthenticateFo, controllers.FOAllUser)         //1
	app.Get("api/mrall", mddlwr.IsAuthenticateMR, controllers.MRAllUser)         //2
	app.Get("api/mrtall", mddlwr.IsAuthenticateMRTI, controllers.MarutiAllUser)  //3
	app.Get("api/rsall", mddlwr.IsAuthenticateRS, controllers.RSAllUser)         //4
	app.Get("api/kkdall", mddlwr.IsAuthenticateKKD, controllers.KiskendaAllUser) //5
	app.Get("api/mliall", mddlwr.IsAuthenticateMLI, controllers.MaliawanAllUser) //6
	app.Get("api/ocall", mddlwr.IsAuthenticateOC, controllers.OCafeAllUser)      //7
	app.Get("api/doall", mddlwr.IsAuthenticateDO, controllers.DOAllUser)         //8
	app.Get("api/vlaall", mddlwr.IsAuthenticateVilla, controllers.VillaAllUser)  //9
	app.Get("api/krkall", mddlwr.IsAuthenticateKR, controllers.KaraokeAllUser)   //10
	//------------------------- 2 Outside Location -------------------------------->
	app.Get("/api/pball", mddlwr.IsAuthenticatePB, controllers.PBAllUser)    //11
	app.Get("/api/tosall", mddlwr.IsAuthenticateTOS, controllers.TOSAllUser) //12
	//------------------------End All User Route Area------------------------------>

	//------------------------------------>
	//--------------------------------------------->
	//----------------------------------------------------------->
	//--------------------------------Admin------------------------------>
	app.Post("/api/createadmin", controllers.RegisterAdmin)
	app.Post("/api/loginadmin", controllers.LoginAdmin)
	//----------Protected------------//
	//----------------------------Set Status Route --------------------------------->
	//--------------------------10 inside location----------------------------->
	app.Put("/api/fosetstatus", mddlwr.IsAuthenticateAdmin, controllers.FOSetStatus)     //1
	app.Put("/api/mrsetstatus", mddlwr.IsAuthenticateAdmin, controllers.MRSetStatus)     //2
	app.Put("/api/mrtsetstatus", mddlwr.IsAuthenticateAdmin, controllers.MRTISetStatus)  //3
	app.Put("/api/rssetstatus", mddlwr.IsAuthenticateAdmin, controllers.RSSetStatus)     //4
	app.Put("/api/kkdsetstatus", mddlwr.IsAuthenticateAdmin, controllers.KKDSetStatus)   //5
	app.Put("/api/mlisetstatus", mddlwr.IsAuthenticateAdmin, controllers.MLISetStatus)   //6
	app.Put("/api/ocsetstatus", mddlwr.IsAuthenticateAdmin, controllers.OCSetStatus)     //7
	app.Put("/api/dosetstatus", mddlwr.IsAuthenticateAdmin, controllers.DOSetStatus)     //8
	app.Put("/api/vlasetstatus", mddlwr.IsAuthenticateAdmin, controllers.VillaSetStatus) //9
	app.Put("/api/krksetstatus", mddlwr.IsAuthenticateAdmin, controllers.KRKSetStatus)   //10
	//----------------------------------------------------------------------------->
	//------------------------- 2 Outside Location -------------------------------->
	app.Put("/api/pbsetstatus", controllers.PBSetStatus)   //11
	app.Put("/api/tossetstatus", controllers.TOSSetStatus) //12
	//------------------------End Set Status User Route Area------------------------------>
	//----------------------------Get All Route --------------------------------->
	//--------------------------10 inside location----------------------------->
	//app.Get("/api/adminalluser", mddlwr.IsAuthenticateAdmin, controllers.AdminGetAll)
	app.Get("/api/fogetall", mddlwr.IsAuthenticateAdmin, controllers.FOGetAll)     //1
	app.Get("/api/mrgetall", mddlwr.IsAuthenticateAdmin, controllers.MRGetAll)     //2
	app.Get("/api/mrtigetall", mddlwr.IsAuthenticateAdmin, controllers.MRTIGetAll) //3
	app.Get("/api/rmgetall", mddlwr.IsAuthenticateAdmin, controllers.RMGetAll)     //4
	app.Get("/api/kkdgetall", mddlwr.IsAuthenticateAdmin, controllers.KKDGetAll)   //5
	app.Get("/api/mligetall", mddlwr.IsAuthenticateAdmin, controllers.MLGetAll)    //6
	app.Get("/api/ocgetall", mddlwr.IsAuthenticateAdmin, controllers.OCGetAll)     //7
	app.Get("/api/dogetall", mddlwr.IsAuthenticateAdmin, controllers.DOGetAll)     //8
	app.Get("/api/vlagetall", mddlwr.IsAuthenticateAdmin, controllers.VillaGetAll) //9
	app.Get("/api/krkgetall", mddlwr.IsAuthenticateAdmin, controllers.KRKGetAll)   //10
	//--------------------------2 outside location----------------------------->
	app.Get("/api/pbgetall", mddlwr.IsAuthenticateAdmin, controllers.PBGetAll) //11
	app.Get("/api/togetall", mddlwr.IsAuthenticateAdmin, controllers.OMGetAll) //12
	//------------------------End Get All Route Area------------------------------>
	//----------------------------Delete Route --------------------------------->
	//--------------------------10 inside location----------------------------->
	app.Delete("/api/fodelete", mddlwr.IsAuthenticateAdmin, controllers.FODeleteUser)     //1
	app.Delete("/api/mrdelete", mddlwr.IsAuthenticateAdmin, controllers.MRDeleteUser)     //2
	app.Delete("/api/mrtidelete", mddlwr.IsAuthenticateAdmin, controllers.MRTIDeleteUser) //3
	app.Delete("/api/rsdelete", mddlwr.IsAuthenticateAdmin, controllers.RSDeleteUser)     //4
	app.Delete("/api/kkddelete", mddlwr.IsAuthenticateAdmin, controllers.KKDDeleteUser)   //5
	app.Delete("/api/mlidelete", mddlwr.IsAuthenticateAdmin, controllers.MLIDeleteUser)   //6
	app.Delete("/api/ocdelete", mddlwr.IsAuthenticateAdmin, controllers.OCDeleteUser)     //7
	app.Delete("/api/dodelete", mddlwr.IsAuthenticateAdmin, controllers.DODeleteUser)     //8
	app.Delete("/api/vladelete", mddlwr.IsAuthenticateAdmin, controllers.VLADeleteUser)   //9
	app.Delete("/api/krkdelete", mddlwr.IsAuthenticateAdmin, controllers.KRKDeleteUser)   //10
	//--------------------------2 outside location----------------------------->
	app.Delete("/api/pbdelete", mddlwr.IsAuthenticateAdmin, controllers.PBDeleteUser)   //11
	app.Delete("/api/tosdelete", mddlwr.IsAuthenticateAdmin, controllers.TOSDeleteUser) //12
	//------------------------End Delete Route Area------------------------------>
	//----------------------------Delete Cookie Route --------------------------------->
	//--------------------------10 inside location----------------------------->
	app.Delete("/api/fodelck", controllers.FODeleteCookie)     //1
	app.Delete("/api/mrdelck", controllers.MRDeleteCookie)     //2
	app.Delete("/api/mrtidelck", controllers.MRTIDeleteCookie) //3
	app.Delete("/api/rsdelck", controllers.RSDeleteCookie)     //4
	app.Delete("/api/kkddelck", controllers.KKDDeleteCookie)   //5
	app.Delete("/api/mlidelck", controllers.MLIDeleteCookie)   //6
	app.Delete("/api/ocdelck", controllers.OCDeleteCookie)     //7
	app.Delete("/api/dodelck", controllers.DODeleteCookie)     //8
	app.Delete("/api/vladelck", controllers.VLADeleteCookie)   //9
	app.Delete("/api/krkdelck", controllers.KRKDeleteCookie)   //10
	//--------------------------2 outside location----------------------------->
	app.Delete("/api/pbdelck", controllers.PBDeleteCookie)   //11
	app.Delete("/api/tosdelck", controllers.TOSDeleteCookie) //12

	app.Delete("/api/admin/logout", mddlwr.IsAuthenticateAdmin, controllers.AdminLogout)
	//--------------------------------------------->
	//---------------------------------->
	//------------------------->
	//------------>

	app.Get("api/gethistory", mddlwr.IsAuthenticateAdmin, controllers.GetHistory)

	////////////////////////////////////PROJECT 2////////////////////////////////////////////////
	app.Post("/api/createsales", salescontroller.RegisterSales)
	app.Put("/api/sales/:id", salescontroller.UpdateSales)
	app.Get("/api/sales/get/:username", salescontroller.GetSales)
	app.Delete("/api/salesdelete/:id", salescontroller.DeleteSales)
	app.Post("/api/upload-avatar", salescontroller.UploadImage)
	app.Static("/api/uploads", "./data/images/avatar")

}
