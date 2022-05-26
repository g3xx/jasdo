document.addEventListener('alpine:init', () => {
    Alpine.store('csrf', {
        val: document.querySelector('meta[name="_csrf"]').content,
    })

    Alpine.store('utils', {
        isLoading: false,
        form : false,
        show: false,
        status: '',
        text : '',

        doMsg(data){
            setTimeout(() => {
                this.show = true;
                this.text = data.Text;
                this.status = data.Status;
            }, 1000);
            if(!this.status.includes("error")){
                setTimeout(() => {
                    this.show = false;
                    this.isLoading = false
                }, 2000);
            }
        },
    })

})

//define function user for page users
function Users() {
    return {
        UserData: {
            Id: null,
			FullName: '',
			Email: '',
			Telp: '',
            Password:'',
            RolesId : {
                Id:''
            }
		},
        mode: "",
        ListUsers : null,
        ListRoles : null,
        pager:null,
        curpage: 1,
        help: Alpine.store("utils"),
        
        test(){
            this.help.form = !this.help.form
        },

        setUser(u){
            this.help.form= !this.help.form
            this.getRoles()
            if (u !== undefined){
                this.mode = "Update User "+ u.Email
                this.UserData.Id = Number(u.Id)
                this.UserData.FullName = u.FullName
                this.UserData.Telp = u.Telp
                this.UserData.Email = u.Email
                this.UserData.Password = ""
                this.UserData.RolesId.Id = Number(u.RolesId.Id)
            }else{
                //this.mode = "Update User "+ u.Email
                this.mode = "Create new User"
                this.UserData.FullName = ""
                this.UserData.Telp = ""
                this.UserData.Email = ""
                this.UserData.Password = ""
                this.UserData.RolesId.Id = 1
            }

        },
        updateOrAddUsers(){
            if(this.mode.includes("new")){
                doPost('/admin/users/', 'POST', this.UserData)
                .then(res => {
                    this.help.doMsg(res)
                }).catch(err => {
                    console.log(err)
                })
            }else {
                id = this.UserData.Id
                doPost('/admin/users/'+id, 'PUT', this.UserData)
                .then(res => {
                    this.help.doMsg(res)
                }).catch(err => {
                    console.log(err)
                })
            }
            setTimeout(() => {
                this.help.form = !this.help.form
                this.getAll() // refest
            }, 1500);
        },

        getAll(e = 0){
            e = e*10-10;
            fetch('/admin/users/?q=all&offset='+e, {
                method: 'GET',
            }).then(res => res.json())
            .then(data => {
                this.ListUsers = data.Data
                if ( data.Total > 10 ){
                    this.pager = pages(this.curpage, data.Total)
                }
            })
        },


        delUsers(e){
            doPost('/admin/users/'+e, 'DELETE')
            .then(res => {
                this.help.doMsg(res)
                location.reload();
            }).catch(err => {
                console.log(err)
            }) 
		
        }, 

        changepage(e){
           this.curpage = e
           this.getAll(e)
        },

        getRoles(){
            fetch('/admin/roles/?q=all', {
                method: 'GET',
            }).then(res => res.json())
            .then(data => {
                this.ListRoles = data
            })

        }

    }
}

function Login(){

    return {
        user : {
            Email : null,
            Pass  : null
        },
        valid : false,
        help : Alpine.store("utils"),
        spin : "off",

        Valid(u){
            if (validateEmail(u.Email)) {
                this.valid = false
            }else{
                this.valid = true
            }
        },

        Auth(){
            this.spin = "on"
            console.log(this.user)
            if(!this.valid){
                doPost('/admin/login', 'POST', this.user)
                .then(res => {
                    if (res.Status.includes("success")){
                        window.location.href = '/admin/users';
                    }else{
                        this.help.doMsg(res)
                    }
                }).catch(err => {
                    console.log(err)
                })    
                setTimeout(() => { this.spin = "off" }, 900);
            }
        }
    }

}