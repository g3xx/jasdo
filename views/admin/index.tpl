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
    <div class="content d-flex flex-column flex-column-fluid">
        <div class="container-xxl">
            <div class="card" x-init="getAll()">
                <div class="card-body p-5 px-lg-19 py-lg-16">

                    <template x-if="ListUsers">
                        <table id="kt_datatable_example_1" class="table table-row-bordered gy-5">
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
                                        <td x-text="user.FullName"></td>
                                        <td x-text="user.Email"></td>
                                        <td x-text="user.Telp"></td>
                                        <td x-text="user.RolesId.Name"></td>
                                        <td>
                                         
                                            
                                            <i class="bi bi-pencil-fill fs-2x text-primary"></i>
                                            <i class="bi bi-trash-fill fs-2x text-danger"></i>
                                        </td>
                                    </tr>
                                </template>
                            </tbody>
                        </table>
                    </template>



                    <a href="/admin/logout/">Logout</a>
                </div>
            </div>
        </div>
    </div>
    <!--end::Content-->

	
		<!--end::Container-->

	