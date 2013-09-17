package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var regionselect_code cu.Function

type regionselect_args struct {
	arg_dst     unsafe.Pointer
	arg_src     unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_region  byte
	arg_N       int
	argptr      [5]unsafe.Pointer
}

// Wrapper for regionselect CUDA kernel, asynchronous.
func k_regionselect_async(dst unsafe.Pointer, src unsafe.Pointer, regions unsafe.Pointer, region byte, N int, cfg *config, str cu.Stream) {
	if regionselect_code == 0 {
		regionselect_code = fatbinLoad(regionselect_map, "regionselect")
	}

	var _a_ regionselect_args

	_a_.arg_dst = dst
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_dst)
	_a_.arg_src = src
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_src)
	_a_.arg_regions = regions
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_regions)
	_a_.arg_region = region
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_region)
	_a_.arg_N = N
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_N)

	args := _a_.argptr[:]
	cu.LaunchKernel(regionselect_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for regionselect CUDA kernel, synchronized.
func k_regionselect(dst unsafe.Pointer, src unsafe.Pointer, regions unsafe.Pointer, region byte, N int, cfg *config) {
	str := stream()
	k_regionselect_async(dst, src, regions, region, N, cfg, str)
	syncAndRecycle(str)
}

var regionselect_map = map[int]string{0: "",
	20: regionselect_ptx_20,
	30: regionselect_ptx_30,
	35: regionselect_ptx_35}

const (
	regionselect_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry regionselect(
	.param .u64 regionselect_param_0,
	.param .u64 regionselect_param_1,
	.param .u64 regionselect_param_2,
	.param .u8 regionselect_param_3,
	.param .u32 regionselect_param_4
)
{
	.reg .pred 	%p<3>;
	.reg .s16 	%rc<3>;
	.reg .s32 	%r<12>;
	.reg .f32 	%f<5>;
	.reg .s64 	%rd<13>;


	ld.param.u64 	%rd5, [regionselect_param_0];
	ld.param.u64 	%rd6, [regionselect_param_1];
	ld.param.u64 	%rd7, [regionselect_param_2];
	ld.param.u8 	%rc1, [regionselect_param_3];
	ld.param.u32 	%r2, [regionselect_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd7;
	.loc 2 5 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 6 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_4;

	.loc 2 7 1
	cvt.s64.s32 	%rd4, %r1;
	add.s64 	%rd8, %rd3, %rd4;
	ld.global.u8 	%rc2, [%rd8];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc2;
	cvt.s16.s8 	%temp2, %rc1;
	setp.ne.s16 	%p2, %temp1, %temp2;
	}
	mov.f32 	%f4, 0f00000000;
	.loc 2 7 1
	@%p2 bra 	BB0_3;

	shl.b64 	%rd9, %rd4, 2;
	add.s64 	%rd10, %rd2, %rd9;
	ld.global.f32 	%f4, [%rd10];

BB0_3:
	shl.b64 	%rd11, %rd4, 2;
	add.s64 	%rd12, %rd1, %rd11;
	st.global.f32 	[%rd12], %f4;

BB0_4:
	.loc 2 9 2
	ret;
}


`
	regionselect_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry regionselect(
	.param .u64 regionselect_param_0,
	.param .u64 regionselect_param_1,
	.param .u64 regionselect_param_2,
	.param .u8 regionselect_param_3,
	.param .u32 regionselect_param_4
)
{
	.reg .pred 	%p<3>;
	.reg .s16 	%rc<3>;
	.reg .s32 	%r<12>;
	.reg .f32 	%f<5>;
	.reg .s64 	%rd<13>;


	ld.param.u64 	%rd5, [regionselect_param_0];
	ld.param.u64 	%rd6, [regionselect_param_1];
	ld.param.u64 	%rd7, [regionselect_param_2];
	ld.param.u8 	%rc1, [regionselect_param_3];
	ld.param.u32 	%r2, [regionselect_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd7;
	.loc 2 5 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 6 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_4;

	.loc 2 7 1
	cvt.s64.s32 	%rd4, %r1;
	add.s64 	%rd8, %rd3, %rd4;
	ld.global.u8 	%rc2, [%rd8];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc2;
	cvt.s16.s8 	%temp2, %rc1;
	setp.ne.s16 	%p2, %temp1, %temp2;
	}
	mov.f32 	%f4, 0f00000000;
	.loc 2 7 1
	@%p2 bra 	BB0_3;

	shl.b64 	%rd9, %rd4, 2;
	add.s64 	%rd10, %rd2, %rd9;
	ld.global.f32 	%f4, [%rd10];

BB0_3:
	shl.b64 	%rd11, %rd4, 2;
	add.s64 	%rd12, %rd1, %rd11;
	st.global.f32 	[%rd12], %f4;

BB0_4:
	.loc 2 9 2
	ret;
}


`
	regionselect_ptx_35 = `
.version 3.1
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
	.loc 2 66 3
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
	.loc 2 71 3
	ret;
}

.visible .entry regionselect(
	.param .u64 regionselect_param_0,
	.param .u64 regionselect_param_1,
	.param .u64 regionselect_param_2,
	.param .u8 regionselect_param_3,
	.param .u32 regionselect_param_4
)
{
	.reg .pred 	%p<3>;
	.reg .s16 	%rc<3>;
	.reg .s32 	%r<11>;
	.reg .f32 	%f<5>;
	.reg .s64 	%rd<13>;


	ld.param.u64 	%rd5, [regionselect_param_0];
	ld.param.u64 	%rd6, [regionselect_param_1];
	ld.param.u64 	%rd7, [regionselect_param_2];
	ld.param.u8 	%rc1, [regionselect_param_3];
	ld.param.u32 	%r2, [regionselect_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd7;
	.loc 3 5 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 3 6 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB2_4;

	.loc 3 7 1
	cvt.s64.s32 	%rd4, %r1;
	add.s64 	%rd8, %rd3, %rd4;
	ld.global.u8 	%rc2, [%rd8];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc2;
	cvt.s16.s8 	%temp2, %rc1;
	setp.ne.s16 	%p2, %temp1, %temp2;
	}
	mov.f32 	%f4, 0f00000000;
	.loc 3 7 1
	@%p2 bra 	BB2_3;

	shl.b64 	%rd9, %rd4, 2;
	add.s64 	%rd10, %rd2, %rd9;
	ld.global.nc.f32 	%f4, [%rd10];

BB2_3:
	shl.b64 	%rd11, %rd4, 2;
	add.s64 	%rd12, %rd1, %rd11;
	st.global.f32 	[%rd12], %f4;

BB2_4:
	.loc 3 9 2
	ret;
}


`
)
