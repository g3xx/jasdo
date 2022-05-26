<div x-show="help.form">
    <form class="form mb-15 fv-plugins-bootstrap5 fv-plugins-framework" id="updateUsers">
        <h1 class="fw-bolder text-dark mb-9" x-text="mode"></h1>
        <div class="d-flex flex-column mb-5 fv-row">
            <label class="fs-5 fw-bold mb-2">Full Name</label>
            <input class="form-control form-control" x-model="UserData.FullName">
        </div>
        <div class="d-flex flex-column mb-5 fv-row">
            <label class="fs-5 fw-bold mb-2">Telp Number</label>
            <div class="input-group mb-5">
                <span class="input-group-text" id="basic-addon1">+62</span>
                <input type="text" class="form-control"  x-model="UserData.Telp" aria-describedby="basic-addon1"/>
            </div>
        
        </div>
        <div class="d-flex flex-column mb-5 fv-row">
            <label class="fs-5 fw-bold mb-2">Email</label>
            <input class="form-control form-control" x-model="UserData.Email">
        </div>

        <div class="d-flex flex-column mb-5 fv-row" >
            <label class="fs-5 fw-bold mb-2">Roles</label>
            <select class="form-select" aria-label="Select example" x-model.number="UserData.RolesId.Id">
                <option>---</option>
                <template x-for="r in ListRoles">
                    <option :value="r.Id" x-text="r.Name"></option>
                </template>
            </select>
        </div>

        <div class="fv-row mb-5" data-kt-password-meter="true">
            <div class="mb-1">
                <label class="form-label fw-bold fs-6 mb-2">Password</label>
                <div class="position-relative mb-3">
                    <input class="form-control form-control-lg form-control-solid"
                        type="password" placeholder="" autocomplete="off" x-model="UserData.Password"/>

                    <span class="btn btn-sm btn-icon position-absolute translate-middle top-50 end-0 me-n2"
                        data-kt-password-meter-control="visibility">
                        <i class="bi bi-eye-slash fs-2"></i>
                        <i class="bi bi-eye fs-2 d-none"></i>
                    </span>
                </div>

                <div class="d-flex align-items-center mb-3" data-kt-password-meter-control="highlight">
                    <div class="flex-grow-1 bg-secondary bg-active-success rounded h-5px me-2"></div>
                    <div class="flex-grow-1 bg-secondary bg-active-success rounded h-5px me-2"></div>
                    <div class="flex-grow-1 bg-secondary bg-active-success rounded h-5px me-2"></div>
                    <div class="flex-grow-1 bg-secondary bg-active-success rounded h-5px"></div>
                </div>
            </div>
            <div class="text-muted">
                Use 8 or more characters with a mix of letters, numbers & symbols.
            </div>
        </div>
        <a @click ="help.form = !help.form" class="btn bg-secondary">Cancel</a>
        <a class="btn btn-primary" @click="updateOrAddUsers()">
            <span class="indicator-label">Save</span>
            <span class="indicator-progress">Please wait... 
            <span class="spinner-border spinner-border-sm align-middle ms-2"></span></span>
        </a>
    </form>
</div>