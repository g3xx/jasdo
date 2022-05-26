{{template "header.tpl" .}}

<div class="d-flex flex-column flex-root" x-data="Login()">
	<!--begin::Authentication - Sign-in -->
	<div class="d-flex flex-column flex-column-fluid bgi-position-y-bottom position-x-center bgi-no-repeat bgi-size-contain bgi-attachment-fixed" style="background-image: url(assets/media/illustrations/dozzy-1/14.png)">
		<!--begin::Content-->
		<div class="d-flex flex-center flex-column flex-column-fluid p-10 pb-lg-20">
			<!--begin::Logo-->
			<a href="../../demo3/dist/index.html" class="mb-12">
				<h2>Logo here</h2>
			</a>
			<div class="w-lg-500px bg-body rounded shadow-sm p-10 p-lg-15 mx-auto">
				<form class="form w-100" id="kt_sign_in_form" data-kt-redirect-url="/index" >
					<div class="text-center mb-10">
						<h1 class="text-dark mb-3">Sign In to Admin Panel</h1>
							<template x-if="help.show">
								<div class="fw-bold fs-4 p-5 mb-10" :class="help.status == 'error' ? 'text-light-danger bg-danger':'text-light-success bg-success'" >
									<p x-text="help.text"></p>
								</div>
							</template>
					</div>
					<div class="fv-row mb-10 fv-plugins-icon-container">
						<label class="form-label fs-6 fw-bolder text-dark">Email</label>
						<input class="form-control form-control-lg form-control-solid" x-model="user.Email" @keyup="Valid(user)"/>
						<div x-show="valid" class="fv-plugins-message-container invalid-feedback">
							<div>The value is not a valid email address</div>
						</div>
					</div>

					<div class="fv-row mb-10 fv-plugins-icon-container">
						<div class="d-flex flex-stack mb-2">
							<label class="form-label fw-bolder text-dark fs-6 mb-0">Password</label>
							<a href="../../demo3/dist/authentication/layouts/basic/password-reset.html" class="link-primary fs-6 fw-bolder">Forgot Password ?</a>
						</div>
						<input class="form-control form-control-lg form-control-solid" type="password" x-model="user.Pass"/>
						<div class="fv-plugins-message-container invalid-feedback"></div>
					</div>
					<div class="text-center">
						<a id="kt_sign_in_submit" class="btn btn-lg btn-primary w-100 mb-5" :data-kt-indicator="spin" @click="Auth()">
							<span class="indicator-label">Continue</span>
							<span class="indicator-progress">Please wait...
							<span class="spinner-border spinner-border-sm align-middle ms-2"></span></span>
						</a>
					</div>
				</form>
			</div>
		</div>
	</div>
</div>

{{template "footer.tpl" .}}