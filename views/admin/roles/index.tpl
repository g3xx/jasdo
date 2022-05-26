<div class="wrapper d-flex flex-column flex-row-fluid" id="kt_wrapper" x-data="Roles()">
    <div id="kt_header" class="header">
        <div class="container d-flex flex-stack flex-wrap gap-2" id="kt_header_container">
            <div class="page-title d-flex flex-column align-items-start justify-content-center flex-wrap me-lg-2 pb-5 pb-lg-0" data-kt-swapper="true" data-kt-swapper-mode="prepend" data-kt-swapper-parent="{default: '#kt_content_container', lg: '#kt_header_container'}">
                <h1 class="d-flex flex-column text-dark fw-bolder my-0 fs-1">List Roles</h1>
            </div>
        </div>
    </div>
    <!--begin::Content-->
    <div class="content d-flex">
        <div class="container-xxl">
            <div class="card" x-init="getAll()">
                <div class="card-body p-5 px-lg-12 py-lg-8">

                    <template x-if="help.show">
                        <div class="p-3 fs-4 mb-10" :class="help.status == 'error' ? 'text-light-danger bg-danger':'text-light-success bg-success'" >
                            <p x-text="help.text"></p>
                        </div>
                    </template>
                
                    <template x-if="ListRoles">
                        <div>
                            <div class="d-flex flex-row-reverse mb-4">
                                <a @click="setRoles()" class="btn btn-dark">Add New Roles</a>
                            </div>
                               
                        <table id="kt_datatable_example_1" class="table table-row-bordered gy-5">
                            <thead>
                                <tr class="fw-bold fs-6 text-muted">
                                    <th>Roles Name</th>
                                    <th>Action</th>
                                </tr>
                            </thead>
                            <tbody>
                                <template x-for="r in ListRoles">
                                    <tr>
                                        <td x-text="r.Id" class="d-none"></td>
                                        <td x-text="r.Name" x-show="targetId != r.Id"></td>
                                        <td x-show="targetId == r.Id">
                                            <input x-model="RolesData.Name" class="form-control"/>
                                        </td>
                                        <td x-show="targetId == r.Id">
                                            <a @click ="targetId = null" class="btn bg-secondary">Cancel</a>
                                            <a class="btn btn-primary" @click="updateOrAddRoles()">
                                                <span class="indicator-label">Save</span>
                                                <span class="indicator-progress">Please wait... 
                                                <span class="spinner-border spinner-border-sm align-middle ms-2"></span></span>
                                            </a>                                  
                                        </td>
                                        <td x-show="targetId!= r.Id">
                                            <a @click="setRoles(r)">
                                                <i class="bi bi-pencil-fill fs-2x text-primary"></i>
                                            </a>
                                            <a @click="RolesData = r" data-bs-toggle="modal" data-bs-target="#modal_delete">
                                                <i class="bi bi-trash-fill fs-2x text-danger"></i>
                                            </a>
                                        </td>
                                    </tr>
                                </template>
                                <tr x-show="help.form">
                                    {{template "/admin/roles/form.tpl" . }}
                                </tr>
                            </tbody>
                        </table>
                    </div>   
                    </template>

                </div>


                <div class="modal fade" tabindex="-1" id="modal_delete">
                    <div class="modal-dialog">
                        <div class="modal-content">
                            <div class="modal-header">
                                <h3 class="modal-title">Delete Roles </h3>
                                   <div class="btn btn-icon btn-sm btn-active-light-primary ms-2" data-bs-dismiss="modal" aria-label="Close">
                                    <span class="svg-icon svg-icon-2x"></span>
                                </div>
                            </div>
                            <div class="modal-body fs-2">
                                <p>Are u sure want to delete ?</p>
                                <p x-text="RolesData.Name" class="text-info"></p>
                            </div>
                            <div class="modal-footer">
                                <button type="button"  @click="targetId = null" class="btn btn-light" data-bs-dismiss="modal">Close</button>
                                <button type="button" class="btn btn-danger" data-bs-dismiss="modal" @click="delRoles(RolesData.Id)">Delete</button>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>
    <!--end::Content-->

	
		<!--end::Container-->

	