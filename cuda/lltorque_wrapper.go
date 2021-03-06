package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for lltorque kernel
var lltorque_code cu.Function

// Stores the arguments for lltorque kernel invocation
type lltorque_args_t struct {
	arg_tx       unsafe.Pointer
	arg_ty       unsafe.Pointer
	arg_tz       unsafe.Pointer
	arg_mx       unsafe.Pointer
	arg_my       unsafe.Pointer
	arg_mz       unsafe.Pointer
	arg_hx       unsafe.Pointer
	arg_hy       unsafe.Pointer
	arg_hz       unsafe.Pointer
	arg_alphaLUT unsafe.Pointer
	arg_regions  unsafe.Pointer
	arg_N        int
	argptr       [12]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for lltorque kernel invocation
var lltorque_args lltorque_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	lltorque_args.argptr[0] = unsafe.Pointer(&lltorque_args.arg_tx)
	lltorque_args.argptr[1] = unsafe.Pointer(&lltorque_args.arg_ty)
	lltorque_args.argptr[2] = unsafe.Pointer(&lltorque_args.arg_tz)
	lltorque_args.argptr[3] = unsafe.Pointer(&lltorque_args.arg_mx)
	lltorque_args.argptr[4] = unsafe.Pointer(&lltorque_args.arg_my)
	lltorque_args.argptr[5] = unsafe.Pointer(&lltorque_args.arg_mz)
	lltorque_args.argptr[6] = unsafe.Pointer(&lltorque_args.arg_hx)
	lltorque_args.argptr[7] = unsafe.Pointer(&lltorque_args.arg_hy)
	lltorque_args.argptr[8] = unsafe.Pointer(&lltorque_args.arg_hz)
	lltorque_args.argptr[9] = unsafe.Pointer(&lltorque_args.arg_alphaLUT)
	lltorque_args.argptr[10] = unsafe.Pointer(&lltorque_args.arg_regions)
	lltorque_args.argptr[11] = unsafe.Pointer(&lltorque_args.arg_N)
}

// Wrapper for lltorque CUDA kernel, asynchronous.
func k_lltorque_async(tx unsafe.Pointer, ty unsafe.Pointer, tz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, hx unsafe.Pointer, hy unsafe.Pointer, hz unsafe.Pointer, alphaLUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("lltorque")
	}

	lltorque_args.Lock()
	defer lltorque_args.Unlock()

	if lltorque_code == 0 {
		lltorque_code = fatbinLoad(lltorque_map, "lltorque")
	}

	lltorque_args.arg_tx = tx
	lltorque_args.arg_ty = ty
	lltorque_args.arg_tz = tz
	lltorque_args.arg_mx = mx
	lltorque_args.arg_my = my
	lltorque_args.arg_mz = mz
	lltorque_args.arg_hx = hx
	lltorque_args.arg_hy = hy
	lltorque_args.arg_hz = hz
	lltorque_args.arg_alphaLUT = alphaLUT
	lltorque_args.arg_regions = regions
	lltorque_args.arg_N = N

	args := lltorque_args.argptr[:]
	cu.LaunchKernel(lltorque_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("lltorque")
	}
}

// maps compute capability on PTX code for lltorque kernel.
var lltorque_map = map[int]string{0: "",
	20: lltorque_ptx_20,
	30: lltorque_ptx_30,
	35: lltorque_ptx_35,
	50: lltorque_ptx_50,
	52: lltorque_ptx_52}

// lltorque PTX code for various compute capabilities.
const (
	lltorque_ptx_20 = `
.version 4.0
.target sm_20
.address_size 64


.visible .entry lltorque(
	.param .u64 lltorque_param_0,
	.param .u64 lltorque_param_1,
	.param .u64 lltorque_param_2,
	.param .u64 lltorque_param_3,
	.param .u64 lltorque_param_4,
	.param .u64 lltorque_param_5,
	.param .u64 lltorque_param_6,
	.param .u64 lltorque_param_7,
	.param .u64 lltorque_param_8,
	.param .u64 lltorque_param_9,
	.param .u64 lltorque_param_10,
	.param .u32 lltorque_param_11
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<35>;
	.reg .s64 	%rd<38>;


	ld.param.u64 	%rd1, [lltorque_param_0];
	ld.param.u64 	%rd2, [lltorque_param_1];
	ld.param.u64 	%rd3, [lltorque_param_2];
	ld.param.u64 	%rd4, [lltorque_param_3];
	ld.param.u64 	%rd5, [lltorque_param_4];
	ld.param.u64 	%rd6, [lltorque_param_5];
	ld.param.u64 	%rd7, [lltorque_param_6];
	ld.param.u64 	%rd8, [lltorque_param_7];
	ld.param.u64 	%rd9, [lltorque_param_8];
	ld.param.u64 	%rd10, [lltorque_param_9];
	ld.param.u64 	%rd11, [lltorque_param_10];
	ld.param.u32 	%r2, [lltorque_param_11];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd12, %rd3;
	cvta.to.global.u64 	%rd13, %rd2;
	cvta.to.global.u64 	%rd14, %rd1;
	cvta.to.global.u64 	%rd15, %rd10;
	cvta.to.global.u64 	%rd16, %rd11;
	cvta.to.global.u64 	%rd17, %rd9;
	cvta.to.global.u64 	%rd18, %rd8;
	cvta.to.global.u64 	%rd19, %rd7;
	cvta.to.global.u64 	%rd20, %rd6;
	cvta.to.global.u64 	%rd21, %rd4;
	cvt.s64.s32	%rd22, %r1;
	mul.wide.s32 	%rd23, %r1, 4;
	add.s64 	%rd24, %rd21, %rd23;
	cvta.to.global.u64 	%rd25, %rd5;
	add.s64 	%rd26, %rd25, %rd23;
	add.s64 	%rd27, %rd20, %rd23;
	add.s64 	%rd28, %rd19, %rd23;
	add.s64 	%rd29, %rd18, %rd23;
	add.s64 	%rd30, %rd17, %rd23;
	add.s64 	%rd31, %rd16, %rd22;
	ld.global.u8 	%rd32, [%rd31];
	shl.b64 	%rd33, %rd32, 2;
	add.s64 	%rd34, %rd15, %rd33;
	ld.global.f32 	%f1, [%rd30];
	ld.global.f32 	%f2, [%rd26];
	mul.f32 	%f3, %f2, %f1;
	ld.global.f32 	%f4, [%rd29];
	ld.global.f32 	%f5, [%rd27];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.f32 	%f8, [%rd28];
	mul.f32 	%f9, %f5, %f8;
	ld.global.f32 	%f10, [%rd24];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	ld.global.f32 	%f16, [%rd34];
	fma.rn.f32 	%f17, %f16, %f16, 0f3F800000;
	mov.f32 	%f18, 0fBF800000;
	div.rn.f32 	%f19, %f18, %f17;
	mul.f32 	%f20, %f2, %f15;
	mul.f32 	%f21, %f5, %f12;
	sub.f32 	%f22, %f20, %f21;
	mul.f32 	%f23, %f5, %f7;
	mul.f32 	%f24, %f10, %f15;
	sub.f32 	%f25, %f23, %f24;
	mul.f32 	%f26, %f10, %f12;
	mul.f32 	%f27, %f2, %f7;
	sub.f32 	%f28, %f26, %f27;
	fma.rn.f32 	%f29, %f16, %f22, %f7;
	fma.rn.f32 	%f30, %f16, %f25, %f12;
	fma.rn.f32 	%f31, %f16, %f28, %f15;
	mul.f32 	%f32, %f19, %f29;
	mul.f32 	%f33, %f19, %f30;
	mul.f32 	%f34, %f19, %f31;
	add.s64 	%rd35, %rd14, %rd23;
	st.global.f32 	[%rd35], %f32;
	add.s64 	%rd36, %rd13, %rd23;
	st.global.f32 	[%rd36], %f33;
	add.s64 	%rd37, %rd12, %rd23;
	st.global.f32 	[%rd37], %f34;

BB0_2:
	ret;
}


`
	lltorque_ptx_30 = `
.version 4.0
.target sm_30
.address_size 64


.visible .entry lltorque(
	.param .u64 lltorque_param_0,
	.param .u64 lltorque_param_1,
	.param .u64 lltorque_param_2,
	.param .u64 lltorque_param_3,
	.param .u64 lltorque_param_4,
	.param .u64 lltorque_param_5,
	.param .u64 lltorque_param_6,
	.param .u64 lltorque_param_7,
	.param .u64 lltorque_param_8,
	.param .u64 lltorque_param_9,
	.param .u64 lltorque_param_10,
	.param .u32 lltorque_param_11
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<35>;
	.reg .s64 	%rd<38>;


	ld.param.u64 	%rd1, [lltorque_param_0];
	ld.param.u64 	%rd2, [lltorque_param_1];
	ld.param.u64 	%rd3, [lltorque_param_2];
	ld.param.u64 	%rd4, [lltorque_param_3];
	ld.param.u64 	%rd5, [lltorque_param_4];
	ld.param.u64 	%rd6, [lltorque_param_5];
	ld.param.u64 	%rd7, [lltorque_param_6];
	ld.param.u64 	%rd8, [lltorque_param_7];
	ld.param.u64 	%rd9, [lltorque_param_8];
	ld.param.u64 	%rd10, [lltorque_param_9];
	ld.param.u64 	%rd11, [lltorque_param_10];
	ld.param.u32 	%r2, [lltorque_param_11];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd12, %rd3;
	cvta.to.global.u64 	%rd13, %rd2;
	cvta.to.global.u64 	%rd14, %rd1;
	cvta.to.global.u64 	%rd15, %rd10;
	cvta.to.global.u64 	%rd16, %rd11;
	cvta.to.global.u64 	%rd17, %rd9;
	cvta.to.global.u64 	%rd18, %rd8;
	cvta.to.global.u64 	%rd19, %rd7;
	cvta.to.global.u64 	%rd20, %rd6;
	cvta.to.global.u64 	%rd21, %rd4;
	cvt.s64.s32	%rd22, %r1;
	mul.wide.s32 	%rd23, %r1, 4;
	add.s64 	%rd24, %rd21, %rd23;
	cvta.to.global.u64 	%rd25, %rd5;
	add.s64 	%rd26, %rd25, %rd23;
	add.s64 	%rd27, %rd20, %rd23;
	add.s64 	%rd28, %rd19, %rd23;
	add.s64 	%rd29, %rd18, %rd23;
	add.s64 	%rd30, %rd17, %rd23;
	add.s64 	%rd31, %rd16, %rd22;
	ld.global.u8 	%rd32, [%rd31];
	shl.b64 	%rd33, %rd32, 2;
	add.s64 	%rd34, %rd15, %rd33;
	ld.global.f32 	%f1, [%rd30];
	ld.global.f32 	%f2, [%rd26];
	mul.f32 	%f3, %f2, %f1;
	ld.global.f32 	%f4, [%rd29];
	ld.global.f32 	%f5, [%rd27];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.f32 	%f8, [%rd28];
	mul.f32 	%f9, %f5, %f8;
	ld.global.f32 	%f10, [%rd24];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	ld.global.f32 	%f16, [%rd34];
	fma.rn.f32 	%f17, %f16, %f16, 0f3F800000;
	mov.f32 	%f18, 0fBF800000;
	div.rn.f32 	%f19, %f18, %f17;
	mul.f32 	%f20, %f2, %f15;
	mul.f32 	%f21, %f5, %f12;
	sub.f32 	%f22, %f20, %f21;
	mul.f32 	%f23, %f5, %f7;
	mul.f32 	%f24, %f10, %f15;
	sub.f32 	%f25, %f23, %f24;
	mul.f32 	%f26, %f10, %f12;
	mul.f32 	%f27, %f2, %f7;
	sub.f32 	%f28, %f26, %f27;
	fma.rn.f32 	%f29, %f16, %f22, %f7;
	fma.rn.f32 	%f30, %f16, %f25, %f12;
	fma.rn.f32 	%f31, %f16, %f28, %f15;
	mul.f32 	%f32, %f19, %f29;
	mul.f32 	%f33, %f19, %f30;
	mul.f32 	%f34, %f19, %f31;
	add.s64 	%rd35, %rd14, %rd23;
	st.global.f32 	[%rd35], %f32;
	add.s64 	%rd36, %rd13, %rd23;
	st.global.f32 	[%rd36], %f33;
	add.s64 	%rd37, %rd12, %rd23;
	st.global.f32 	[%rd37], %f34;

BB0_2:
	ret;
}


`
	lltorque_ptx_35 = `
.version 4.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry lltorque(
	.param .u64 lltorque_param_0,
	.param .u64 lltorque_param_1,
	.param .u64 lltorque_param_2,
	.param .u64 lltorque_param_3,
	.param .u64 lltorque_param_4,
	.param .u64 lltorque_param_5,
	.param .u64 lltorque_param_6,
	.param .u64 lltorque_param_7,
	.param .u64 lltorque_param_8,
	.param .u64 lltorque_param_9,
	.param .u64 lltorque_param_10,
	.param .u32 lltorque_param_11
)
{
	.reg .pred 	%p<2>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<35>;
	.reg .s64 	%rd<39>;


	ld.param.u64 	%rd1, [lltorque_param_0];
	ld.param.u64 	%rd2, [lltorque_param_1];
	ld.param.u64 	%rd3, [lltorque_param_2];
	ld.param.u64 	%rd4, [lltorque_param_3];
	ld.param.u64 	%rd5, [lltorque_param_4];
	ld.param.u64 	%rd6, [lltorque_param_5];
	ld.param.u64 	%rd7, [lltorque_param_6];
	ld.param.u64 	%rd8, [lltorque_param_7];
	ld.param.u64 	%rd9, [lltorque_param_8];
	ld.param.u64 	%rd10, [lltorque_param_9];
	ld.param.u64 	%rd11, [lltorque_param_10];
	ld.param.u32 	%r2, [lltorque_param_11];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB5_2;

	cvta.to.global.u64 	%rd12, %rd3;
	cvta.to.global.u64 	%rd13, %rd2;
	cvta.to.global.u64 	%rd14, %rd1;
	cvta.to.global.u64 	%rd15, %rd10;
	cvta.to.global.u64 	%rd16, %rd11;
	cvta.to.global.u64 	%rd17, %rd9;
	cvta.to.global.u64 	%rd18, %rd8;
	cvta.to.global.u64 	%rd19, %rd7;
	cvta.to.global.u64 	%rd20, %rd6;
	cvta.to.global.u64 	%rd21, %rd5;
	cvta.to.global.u64 	%rd22, %rd4;
	cvt.s64.s32	%rd23, %r1;
	mul.wide.s32 	%rd24, %r1, 4;
	add.s64 	%rd25, %rd22, %rd24;
	add.s64 	%rd26, %rd21, %rd24;
	add.s64 	%rd27, %rd20, %rd24;
	add.s64 	%rd28, %rd19, %rd24;
	add.s64 	%rd29, %rd18, %rd24;
	add.s64 	%rd30, %rd17, %rd24;
	add.s64 	%rd31, %rd16, %rd23;
	ld.global.nc.u8 	%rs1, [%rd31];
	cvt.u64.u16	%rd32, %rs1;
	and.b64  	%rd33, %rd32, 255;
	shl.b64 	%rd34, %rd33, 2;
	add.s64 	%rd35, %rd15, %rd34;
	ld.global.nc.f32 	%f1, [%rd30];
	ld.global.nc.f32 	%f2, [%rd26];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd29];
	ld.global.nc.f32 	%f5, [%rd27];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd28];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd25];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	ld.global.nc.f32 	%f16, [%rd35];
	fma.rn.f32 	%f17, %f16, %f16, 0f3F800000;
	mov.f32 	%f18, 0fBF800000;
	div.rn.f32 	%f19, %f18, %f17;
	mul.f32 	%f20, %f2, %f15;
	mul.f32 	%f21, %f5, %f12;
	sub.f32 	%f22, %f20, %f21;
	mul.f32 	%f23, %f5, %f7;
	mul.f32 	%f24, %f10, %f15;
	sub.f32 	%f25, %f23, %f24;
	mul.f32 	%f26, %f10, %f12;
	mul.f32 	%f27, %f2, %f7;
	sub.f32 	%f28, %f26, %f27;
	fma.rn.f32 	%f29, %f16, %f22, %f7;
	fma.rn.f32 	%f30, %f16, %f25, %f12;
	fma.rn.f32 	%f31, %f16, %f28, %f15;
	mul.f32 	%f32, %f19, %f29;
	mul.f32 	%f33, %f19, %f30;
	mul.f32 	%f34, %f19, %f31;
	add.s64 	%rd36, %rd14, %rd24;
	st.global.f32 	[%rd36], %f32;
	add.s64 	%rd37, %rd13, %rd24;
	st.global.f32 	[%rd37], %f33;
	add.s64 	%rd38, %rd12, %rd24;
	st.global.f32 	[%rd38], %f34;

BB5_2:
	ret;
}


`
	lltorque_ptx_50 = `
.version 4.2
.target sm_50
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	lltorque
.visible .entry lltorque(
	.param .u64 lltorque_param_0,
	.param .u64 lltorque_param_1,
	.param .u64 lltorque_param_2,
	.param .u64 lltorque_param_3,
	.param .u64 lltorque_param_4,
	.param .u64 lltorque_param_5,
	.param .u64 lltorque_param_6,
	.param .u64 lltorque_param_7,
	.param .u64 lltorque_param_8,
	.param .u64 lltorque_param_9,
	.param .u64 lltorque_param_10,
	.param .u32 lltorque_param_11
)
{
	.reg .pred 	%p<2>;
	.reg .s16 	%rs<2>;
	.reg .f32 	%f<35>;
	.reg .s32 	%r<11>;
	.reg .s64 	%rd<37>;


	ld.param.u64 	%rd1, [lltorque_param_0];
	ld.param.u64 	%rd2, [lltorque_param_1];
	ld.param.u64 	%rd3, [lltorque_param_2];
	ld.param.u64 	%rd4, [lltorque_param_3];
	ld.param.u64 	%rd5, [lltorque_param_4];
	ld.param.u64 	%rd6, [lltorque_param_5];
	ld.param.u64 	%rd7, [lltorque_param_6];
	ld.param.u64 	%rd8, [lltorque_param_7];
	ld.param.u64 	%rd9, [lltorque_param_8];
	ld.param.u64 	%rd10, [lltorque_param_9];
	ld.param.u64 	%rd11, [lltorque_param_10];
	ld.param.u32 	%r2, [lltorque_param_11];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB6_2;

	cvta.to.global.u64 	%rd12, %rd4;
	cvt.s64.s32	%rd13, %r1;
	mul.wide.s32 	%rd14, %r1, 4;
	add.s64 	%rd15, %rd12, %rd14;
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd14;
	cvta.to.global.u64 	%rd18, %rd6;
	add.s64 	%rd19, %rd18, %rd14;
	cvta.to.global.u64 	%rd20, %rd7;
	add.s64 	%rd21, %rd20, %rd14;
	cvta.to.global.u64 	%rd22, %rd8;
	add.s64 	%rd23, %rd22, %rd14;
	cvta.to.global.u64 	%rd24, %rd9;
	add.s64 	%rd25, %rd24, %rd14;
	cvta.to.global.u64 	%rd26, %rd11;
	add.s64 	%rd27, %rd26, %rd13;
	ld.global.nc.u8 	%rs1, [%rd27];
	cvta.to.global.u64 	%rd28, %rd10;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd29, %r10, 4;
	add.s64 	%rd30, %rd28, %rd29;
	ld.global.nc.f32 	%f1, [%rd25];
	ld.global.nc.f32 	%f2, [%rd17];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd23];
	ld.global.nc.f32 	%f5, [%rd19];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd21];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd15];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	ld.global.nc.f32 	%f16, [%rd30];
	fma.rn.f32 	%f17, %f16, %f16, 0f3F800000;
	mov.f32 	%f18, 0fBF800000;
	div.rn.f32 	%f19, %f18, %f17;
	mul.f32 	%f20, %f2, %f15;
	mul.f32 	%f21, %f5, %f12;
	sub.f32 	%f22, %f20, %f21;
	mul.f32 	%f23, %f5, %f7;
	mul.f32 	%f24, %f10, %f15;
	sub.f32 	%f25, %f23, %f24;
	mul.f32 	%f26, %f10, %f12;
	mul.f32 	%f27, %f2, %f7;
	sub.f32 	%f28, %f26, %f27;
	fma.rn.f32 	%f29, %f16, %f22, %f7;
	fma.rn.f32 	%f30, %f16, %f25, %f12;
	fma.rn.f32 	%f31, %f16, %f28, %f15;
	mul.f32 	%f32, %f19, %f29;
	mul.f32 	%f33, %f19, %f30;
	mul.f32 	%f34, %f19, %f31;
	cvta.to.global.u64 	%rd31, %rd1;
	add.s64 	%rd32, %rd31, %rd14;
	st.global.f32 	[%rd32], %f32;
	cvta.to.global.u64 	%rd33, %rd2;
	add.s64 	%rd34, %rd33, %rd14;
	st.global.f32 	[%rd34], %f33;
	cvta.to.global.u64 	%rd35, %rd3;
	add.s64 	%rd36, %rd35, %rd14;
	st.global.f32 	[%rd36], %f34;

BB6_2:
	ret;
}


`
	lltorque_ptx_52 = `
.version 4.2
.target sm_52
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	lltorque
.visible .entry lltorque(
	.param .u64 lltorque_param_0,
	.param .u64 lltorque_param_1,
	.param .u64 lltorque_param_2,
	.param .u64 lltorque_param_3,
	.param .u64 lltorque_param_4,
	.param .u64 lltorque_param_5,
	.param .u64 lltorque_param_6,
	.param .u64 lltorque_param_7,
	.param .u64 lltorque_param_8,
	.param .u64 lltorque_param_9,
	.param .u64 lltorque_param_10,
	.param .u32 lltorque_param_11
)
{
	.reg .pred 	%p<2>;
	.reg .s16 	%rs<2>;
	.reg .f32 	%f<35>;
	.reg .s32 	%r<11>;
	.reg .s64 	%rd<37>;


	ld.param.u64 	%rd1, [lltorque_param_0];
	ld.param.u64 	%rd2, [lltorque_param_1];
	ld.param.u64 	%rd3, [lltorque_param_2];
	ld.param.u64 	%rd4, [lltorque_param_3];
	ld.param.u64 	%rd5, [lltorque_param_4];
	ld.param.u64 	%rd6, [lltorque_param_5];
	ld.param.u64 	%rd7, [lltorque_param_6];
	ld.param.u64 	%rd8, [lltorque_param_7];
	ld.param.u64 	%rd9, [lltorque_param_8];
	ld.param.u64 	%rd10, [lltorque_param_9];
	ld.param.u64 	%rd11, [lltorque_param_10];
	ld.param.u32 	%r2, [lltorque_param_11];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB6_2;

	cvta.to.global.u64 	%rd12, %rd4;
	cvt.s64.s32	%rd13, %r1;
	mul.wide.s32 	%rd14, %r1, 4;
	add.s64 	%rd15, %rd12, %rd14;
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd14;
	cvta.to.global.u64 	%rd18, %rd6;
	add.s64 	%rd19, %rd18, %rd14;
	cvta.to.global.u64 	%rd20, %rd7;
	add.s64 	%rd21, %rd20, %rd14;
	cvta.to.global.u64 	%rd22, %rd8;
	add.s64 	%rd23, %rd22, %rd14;
	cvta.to.global.u64 	%rd24, %rd9;
	add.s64 	%rd25, %rd24, %rd14;
	cvta.to.global.u64 	%rd26, %rd11;
	add.s64 	%rd27, %rd26, %rd13;
	ld.global.nc.u8 	%rs1, [%rd27];
	cvta.to.global.u64 	%rd28, %rd10;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd29, %r10, 4;
	add.s64 	%rd30, %rd28, %rd29;
	ld.global.nc.f32 	%f1, [%rd25];
	ld.global.nc.f32 	%f2, [%rd17];
	mul.f32 	%f3, %f2, %f1;
	ld.global.nc.f32 	%f4, [%rd23];
	ld.global.nc.f32 	%f5, [%rd19];
	mul.f32 	%f6, %f5, %f4;
	sub.f32 	%f7, %f3, %f6;
	ld.global.nc.f32 	%f8, [%rd21];
	mul.f32 	%f9, %f5, %f8;
	ld.global.nc.f32 	%f10, [%rd15];
	mul.f32 	%f11, %f10, %f1;
	sub.f32 	%f12, %f9, %f11;
	mul.f32 	%f13, %f10, %f4;
	mul.f32 	%f14, %f2, %f8;
	sub.f32 	%f15, %f13, %f14;
	ld.global.nc.f32 	%f16, [%rd30];
	fma.rn.f32 	%f17, %f16, %f16, 0f3F800000;
	mov.f32 	%f18, 0fBF800000;
	div.rn.f32 	%f19, %f18, %f17;
	mul.f32 	%f20, %f2, %f15;
	mul.f32 	%f21, %f5, %f12;
	sub.f32 	%f22, %f20, %f21;
	mul.f32 	%f23, %f5, %f7;
	mul.f32 	%f24, %f10, %f15;
	sub.f32 	%f25, %f23, %f24;
	mul.f32 	%f26, %f10, %f12;
	mul.f32 	%f27, %f2, %f7;
	sub.f32 	%f28, %f26, %f27;
	fma.rn.f32 	%f29, %f16, %f22, %f7;
	fma.rn.f32 	%f30, %f16, %f25, %f12;
	fma.rn.f32 	%f31, %f16, %f28, %f15;
	mul.f32 	%f32, %f19, %f29;
	mul.f32 	%f33, %f19, %f30;
	mul.f32 	%f34, %f19, %f31;
	cvta.to.global.u64 	%rd31, %rd1;
	add.s64 	%rd32, %rd31, %rd14;
	st.global.f32 	[%rd32], %f32;
	cvta.to.global.u64 	%rd33, %rd2;
	add.s64 	%rd34, %rd33, %rd14;
	st.global.f32 	[%rd34], %f33;
	cvta.to.global.u64 	%rd35, %rd3;
	add.s64 	%rd36, %rd35, %rd14;
	st.global.f32 	[%rd36], %f34;

BB6_2:
	ret;
}


`
)
