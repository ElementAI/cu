package cu

import "github.com/pkg/errors"

const add32PTX = `//
// Generated by NVIDIA NVVM Compiler
//
// Compiler Build ID: CL-21554848
// Cuda compilation tools, release 8.0, V8.0.61
// Based on LLVM 3.4svn
//

.version 5.0
.target sm_30
.address_size 64

	// .globl	add32

.visible .entry add32(
	.param .u64 add32_param_0,
	.param .u64 add32_param_1,
	.param .u32 add32_param_2
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<4>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<8>;


	ld.param.u64 	%rd1, [add32_param_0];
	ld.param.u64 	%rd2, [add32_param_1];
	ld.param.u32 	%r2, [add32_param_2];
	mov.u32 	%r3, %ctaid.x;
	mov.u32 	%r4, %ctaid.z;
	mov.u32 	%r5, %nctaid.y;
	mov.u32 	%r6, %ctaid.y;
	mad.lo.s32 	%r7, %r4, %r5, %r6;
	mov.u32 	%r8, %nctaid.x;
	mad.lo.s32 	%r9, %r7, %r8, %r3;
	mov.u32 	%r10, %ntid.y;
	mov.u32 	%r11, %ntid.x;
	mul.lo.s32 	%r12, %r10, %r11;
	mov.u32 	%r13, %ntid.z;
	mov.u32 	%r14, %tid.y;
	mov.u32 	%r15, %tid.z;
	mad.lo.s32 	%r16, %r9, %r13, %r15;
	mov.u32 	%r17, %tid.x;
	mad.lo.s32 	%r18, %r14, %r11, %r17;
	mad.lo.s32 	%r1, %r12, %r16, %r18;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd3, %rd1;
	mul.wide.s32 	%rd4, %r1, 4;
	add.s64 	%rd5, %rd3, %rd4;
	cvta.to.global.u64 	%rd6, %rd2;
	add.s64 	%rd7, %rd6, %rd4;
	ld.global.f32 	%f1, [%rd7];
	ld.global.f32 	%f2, [%rd5];
	add.rn.f32 	%f3, %f2, %f1;
	st.global.f32 	[%rd5], %f3;

BB0_2:
	ret;
}


`

func testSetup() (dev Device, ctx CUContext, err error) {
	devices, _ := NumDevices()

	if devices == 0 {
		err = errors.Errorf("NoDevice")
		return
	}

	dev = Device(0)
	if ctx, err = dev.MakeContext(SchedAuto); err != nil {
		return
	}
	return
}

func testTeardown(ctx CUContext, mod Module) {
	if (mod != Module{}) {
		mod.Unload()
	}
	ctx.Destroy()
}
