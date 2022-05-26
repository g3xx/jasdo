

function Roles() {
    return {
        RolesData: { Id: null, Name: ''},
        ListRoles: null,
        help: Alpine.store("utils"),
        targetId:null,
        mode : null,
        
        test(){
            //this.help.form= !this.help.form
        },

        setRoles(u){
            
            if (u !== undefined){
                this.targetId = Number(u.Id)
                this.RolesData.Id = Number(u.Id)
                this.RolesData.Name = u.Name
                this.help.form= false
            }else{
                this.targetId = null
                this.help.form= !this.help.form
                this.RolesData.Id = null
                this.RolesData.Name = ""
            }

        },
        getAll(){
            fetch('/admin/roles/?q=all', {
                method: 'GET',
            }).then(res => res.json())
            .then(data => {
                this.ListRoles = data
            })
        },

        updateOrAddRoles(){
            console.log(this.RolesData)

            this.help.form = false
            if (this.targetId){
                id = this.RolesData.Id
                doPost('/admin/roles/'+id, 'PUT', this.RolesData)
                .then(res => {
                    this.help.doMsg(res)
      
                    this.targetId = null
                }).catch(err => {
                    console.log(err)
                })
            }else{
                doPost('/admin/roles/', 'POST', this.RolesData)
                .then(res => {
                    this.help.doMsg(res)            

                }).catch(err => {
                    console.log(err)
                })
            }
            setTimeout(() => {
                this.getAll()
            }, 500);

        },

        delRoles(e){
            doPost('/admin/roles/'+e, 'DELETE')
            .then(res => {
                this.help.doMsg(res)
                setTimeout(() => {
                    this.getAll()
                }, 500);
            }).catch(err => {
                console.log(err)
            }) 
		
        }, 
    }

}