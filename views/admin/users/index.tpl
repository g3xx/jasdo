<div class="wrapper d-flex flex-column flex-row-fluid" id="kt_wrapper" x-data="Users()">
	<!--begin::Header-->
    <div id="kt_header" class="header">
        <!--begin::Container-->
        <div class="container d-flex flex-stack flex-wrap gap-2" id="kt_header_container">
            <!--begin::Page title-->
            <div class="page-title d-flex flex-column align-items-start justify-content-center flex-wrap me-lg-2 pb-5 pb-lg-0" data-kt-swapper="true" data-kt-swapper-mode="prepend" data-kt-swapper-parent="{default: '#kt_content_container', lg: '#kt_header_container'}">
                <!--begin::Heading-->
                <h1 class="d-flex flex-column text-dark fw-bolder my-0 fs-1">List Users</h1>
                <!--end::Heading-->
            </div>
        </div>
    </div>
    <!--begin::Content-->
    <div class="content d-flex">
        <div class="container-xxl">
            <div class="card" x-init="getAll()">
                <div class="card-body p-5 px-lg-12 py-lg-8">

                    <template x-if="help.show" >
                        <div class="p-3 fs-4 mb-10" :class="help.status == 'error' ? 'text-light-danger bg-danger':'text-light-success bg-success'" >
                            <p x-text="help.text"></p>
                        </div>
                    </template>
                
                    <template x-if="ListUsers">
                        <div x-show="!help.form">

                        <div class="d-flex flex-row-reverse mb-4">
                            <a @click="setUser()" class="btn btn-dark">Add User</a>
                        </div>
                        <table x-show="!help.form" id="kt_datatable_example_1" class="table table-row-bordered gy-5">
                            <thead>
                                <tr class="fw-bold fs-6 text-muted">
                                    <th>Full Name</th>
                                    <th>Email</th>
                                    <th>Telp</th>
                                    <th>Roles</th>
                                    <th>Action</th>
                                </tr>
                            </thead>
                            <tbody>
                                <template x-for="user in ListUsers">
                                    <tr>
                                        <td x-text="user.Id" class="d-none"></td>
                                        <td x-text="user.FullName"></td>
                                        <td x-text="user.Email"></td>
                                        <td x-text="user.Telp"></td>
                                        <td x-text="user.RolesId.Name"></td>
                                        <td>
                                            <a @click="setUser(user)">
                                                <i class="bi bi-pencil-fill fs-2x text-primary"></i>
                                            </a>
                                            <a @click="UserData = user" data-bs-toggle="modal" data-bs-target="#modal_delete">
                                                <i class="bi bi-trash-fill fs-2x text-danger"></i>
                                            </a>
                                        </td>
                                    </tr>
                                </template>
                            </tbody>
                        </table>

                            <div x-show="pager" class="d-flex justify-content-center">
                                <ul class="pagination">
                                    <li class="page-item previous disabled"><a href="#" class="page-link"><i class="previous"></i></a></li>
                                    <template x-for="i in pager">
                                        
                                            <li class="page-item" :class="curpage == i ? 'active' : ''">
                                                <template x-if="i != '...'">
                                                    <a @click="changepage(i)" class="page-link" x-text="i"></a>
                                                </template>
                                                <div x-show="i =='...'">
                                                    <a class="page-link" x-text="i"></a>
                                                </div>
                                            </li>
                                    </template>
                                    <li class="page-item next"><a @click="changepage(curpage+1)" class="page-link"><i class="next"></i></a></li>
                                </ul>
                            </div>
                        </div>
                        
                    </template>

                    {{template "/admin/users/form.tpl" . }}

                </div>


                <div class="modal fade" tabindex="-1" id="modal_delete">
                    <div class="modal-dialog">
                        <div class="modal-content">
                            <div class="modal-header">
                                <h3 class="modal-title">Delete Users</h3>
                                   <div class="btn btn-icon btn-sm btn-active-light-primary ms-2" data-bs-dismiss="modal" aria-label="Close">
                                    <span class="svg-icon svg-icon-2x"></span>
                                </div>
                            </div>
                            <div class="modal-body fs-2">
                                <p>Are u sure want to delete ?</p>
                                <p x-text="UserData.Email" class="text-info"></p>
                            </div>
                            <div class="modal-footer">
                                <button type="button" class="btn btn-light" data-bs-dismiss="modal">Close</button>
                                <button type="button" class="btn btn-danger" data-bs-dismiss="modal" @click="delUsers(UserData.Id)">Delete</button>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>
    <!--end::Content-->

	
		<!--end::Container-->

	