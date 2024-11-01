//go:build amd64
// +build amd64

// Code generated by asm2asm, DO NOT EDIT.

package sse

var _text_validate_utf8_fast = []byte{
	// .p2align 4, 0x90
	// _validate_utf8_fast
	0x55,             // pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000001 movq         %rsp, %rbp
	0x53,             //0x00000004 pushq        %rbx
	0x50,             //0x00000005 pushq        %rax
	0x4c, 0x8b, 0x17, //0x00000006 movq         (%rdi), %r10
	0x4c, 0x8b, 0x5f, 0x08, //0x00000009 movq         $8(%rdi), %r11
	0x4b, 0x8d, 0x34, 0x1a, //0x0000000d leaq         (%r10,%r11), %rsi
	0x48, 0x83, 0xc6, 0xfd, //0x00000011 addq         $-3, %rsi
	0x4c, 0x89, 0xd0, //0x00000015 movq         %r10, %rax
	0x4c, 0x39, 0xd6, //0x00000018 cmpq         %r10, %rsi
	0x0f, 0x86, 0xdd, 0x00, 0x00, 0x00, //0x0000001b jbe          LBB0_14
	0x4c, 0x89, 0xd0, //0x00000021 movq         %r10, %rax
	0xe9, 0x13, 0x00, 0x00, 0x00, //0x00000024 jmp          LBB0_3
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000029 .p2align 4, 0x90
	//0x00000030 LBB0_2
	0x48, 0x01, 0xd0, //0x00000030 addq         %rdx, %rax
	0x48, 0x39, 0xf0, //0x00000033 cmpq         %rsi, %rax
	0x0f, 0x83, 0xc2, 0x00, 0x00, 0x00, //0x00000036 jae          LBB0_14
	//0x0000003c LBB0_3
	0xba, 0x01, 0x00, 0x00, 0x00, //0x0000003c movl         $1, %edx
	0x80, 0x38, 0x00, //0x00000041 cmpb         $0, (%rax)
	0x0f, 0x89, 0xe6, 0xff, 0xff, 0xff, //0x00000044 jns          LBB0_2
	0x8b, 0x38, //0x0000004a movl         (%rax), %edi
	0x89, 0xf9, //0x0000004c movl         %edi, %ecx
	0x81, 0xe1, 0xf0, 0xc0, 0xc0, 0x00, //0x0000004e andl         $12632304, %ecx
	0x81, 0xf9, 0xe0, 0x80, 0x80, 0x00, //0x00000054 cmpl         $8421600, %ecx
	0x0f, 0x85, 0x30, 0x00, 0x00, 0x00, //0x0000005a jne          LBB0_7
	0x89, 0xf9, //0x00000060 movl         %edi, %ecx
	0x81, 0xe1, 0x0f, 0x20, 0x00, 0x00, //0x00000062 andl         $8207, %ecx
	0x81, 0xf9, 0x0d, 0x20, 0x00, 0x00, //0x00000068 cmpl         $8205, %ecx
	0x0f, 0x84, 0x1c, 0x00, 0x00, 0x00, //0x0000006e je           LBB0_7
	0xba, 0x03, 0x00, 0x00, 0x00, //0x00000074 movl         $3, %edx
	0x85, 0xc9, //0x00000079 testl        %ecx, %ecx
	0x0f, 0x85, 0xaf, 0xff, 0xff, 0xff, //0x0000007b jne          LBB0_2
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000081 .p2align 4, 0x90
	//0x00000090 LBB0_7
	0x89, 0xf9, //0x00000090 movl         %edi, %ecx
	0x81, 0xe1, 0xe0, 0xc0, 0x00, 0x00, //0x00000092 andl         $49376, %ecx
	0x81, 0xf9, 0xc0, 0x80, 0x00, 0x00, //0x00000098 cmpl         $32960, %ecx
	0x0f, 0x85, 0x10, 0x00, 0x00, 0x00, //0x0000009e jne          LBB0_9
	0x89, 0xf9, //0x000000a4 movl         %edi, %ecx
	0xba, 0x02, 0x00, 0x00, 0x00, //0x000000a6 movl         $2, %edx
	0x83, 0xe1, 0x1e, //0x000000ab andl         $30, %ecx
	0x0f, 0x85, 0x7c, 0xff, 0xff, 0xff, //0x000000ae jne          LBB0_2
	//0x000000b4 LBB0_9
	0x89, 0xf9, //0x000000b4 movl         %edi, %ecx
	0x81, 0xe1, 0xf8, 0xc0, 0xc0, 0xc0, //0x000000b6 andl         $-1061109512, %ecx
	0x81, 0xf9, 0xf0, 0x80, 0x80, 0x80, //0x000000bc cmpl         $-2139062032, %ecx
	0x0f, 0x85, 0x29, 0x00, 0x00, 0x00, //0x000000c2 jne          LBB0_13
	0x89, 0xf9, //0x000000c8 movl         %edi, %ecx
	0x81, 0xe1, 0x07, 0x30, 0x00, 0x00, //0x000000ca andl         $12295, %ecx
	0x0f, 0x84, 0x1b, 0x00, 0x00, 0x00, //0x000000d0 je           LBB0_13
	0xba, 0x04, 0x00, 0x00, 0x00, //0x000000d6 movl         $4, %edx
	0x40, 0xf6, 0xc7, 0x04, //0x000000db testb        $4, %dil
	0x0f, 0x84, 0x4b, 0xff, 0xff, 0xff, //0x000000df je           LBB0_2
	0x81, 0xe7, 0x03, 0x30, 0x00, 0x00, //0x000000e5 andl         $12291, %edi
	0x0f, 0x84, 0x3f, 0xff, 0xff, 0xff, //0x000000eb je           LBB0_2
	//0x000000f1 LBB0_13
	0x48, 0xf7, 0xd0, //0x000000f1 notq         %rax
	0x4c, 0x01, 0xd0, //0x000000f4 addq         %r10, %rax
	0x48, 0x83, 0xc4, 0x08, //0x000000f7 addq         $8, %rsp
	0x5b, //0x000000fb popq         %rbx
	0x5d, //0x000000fc popq         %rbp
	0xc3, //0x000000fd retq
	//0x000000fe LBB0_14
	0x4d, 0x01, 0xd3, //0x000000fe addq         %r10, %r11
	0x4c, 0x39, 0xd8, //0x00000101 cmpq         %r11, %rax
	0x0f, 0x83, 0x03, 0x01, 0x00, 0x00, //0x00000104 jae          LBB0_30
	0x4c, 0x8d, 0x45, 0xf4, //0x0000010a leaq         $-12(%rbp), %r8
	0x4c, 0x8d, 0x4d, 0xf2, //0x0000010e leaq         $-14(%rbp), %r9
	0xe9, 0x16, 0x00, 0x00, 0x00, //0x00000112 jmp          LBB0_17
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000117 .p2align 4, 0x90
	//0x00000120 LBB0_16
	0x48, 0x83, 0xc0, 0x01, //0x00000120 addq         $1, %rax
	0x4c, 0x39, 0xd8, //0x00000124 cmpq         %r11, %rax
	0x0f, 0x83, 0xe0, 0x00, 0x00, 0x00, //0x00000127 jae          LBB0_30
	//0x0000012d LBB0_17
	0x80, 0x38, 0x00, //0x0000012d cmpb         $0, (%rax)
	0x0f, 0x89, 0xea, 0xff, 0xff, 0xff, //0x00000130 jns          LBB0_16
	0xc6, 0x45, 0xf4, 0x00, //0x00000136 movb         $0, $-12(%rbp)
	0xc6, 0x45, 0xf2, 0x00, //0x0000013a movb         $0, $-14(%rbp)
	0x4c, 0x89, 0xda, //0x0000013e movq         %r11, %rdx
	0x48, 0x29, 0xc2, //0x00000141 subq         %rax, %rdx
	0x48, 0x83, 0xfa, 0x02, //0x00000144 cmpq         $2, %rdx
	0x0f, 0x82, 0x31, 0x00, 0x00, 0x00, //0x00000148 jb           LBB0_21
	0x0f, 0xb6, 0x30, //0x0000014e movzbl       (%rax), %esi
	0x0f, 0xb6, 0x78, 0x01, //0x00000151 movzbl       $1(%rax), %edi
	0x40, 0x88, 0x75, 0xf4, //0x00000155 movb         %sil, $-12(%rbp)
	0x48, 0x8d, 0x48, 0x02, //0x00000159 leaq         $2(%rax), %rcx
	0x48, 0x83, 0xc2, 0xfe, //0x0000015d addq         $-2, %rdx
	0x4c, 0x89, 0xcb, //0x00000161 movq         %r9, %rbx
	0x48, 0x85, 0xd2, //0x00000164 testq        %rdx, %rdx
	0x0f, 0x84, 0x25, 0x00, 0x00, 0x00, //0x00000167 je           LBB0_22
	//0x0000016d LBB0_20
	0x0f, 0xb6, 0x09, //0x0000016d movzbl       (%rcx), %ecx
	0x88, 0x0b, //0x00000170 movb         %cl, (%rbx)
	0x0f, 0xb6, 0x75, 0xf4, //0x00000172 movzbl       $-12(%rbp), %esi
	0x0f, 0xb6, 0x4d, 0xf2, //0x00000176 movzbl       $-14(%rbp), %ecx
	0xe9, 0x15, 0x00, 0x00, 0x00, //0x0000017a jmp          LBB0_23
	//0x0000017f LBB0_21
	0x31, 0xf6, //0x0000017f xorl         %esi, %esi
	0x31, 0xff, //0x00000181 xorl         %edi, %edi
	0x4c, 0x89, 0xc3, //0x00000183 movq         %r8, %rbx
	0x48, 0x89, 0xc1, //0x00000186 movq         %rax, %rcx
	0x48, 0x85, 0xd2, //0x00000189 testq        %rdx, %rdx
	0x0f, 0x85, 0xdb, 0xff, 0xff, 0xff, //0x0000018c jne          LBB0_20
	//0x00000192 LBB0_22
	0x31, 0xc9, //0x00000192 xorl         %ecx, %ecx
	//0x00000194 LBB0_23
	0x0f, 0xb6, 0xc9, //0x00000194 movzbl       %cl, %ecx
	0xc1, 0xe1, 0x10, //0x00000197 shll         $16, %ecx
	0x40, 0x0f, 0xb6, 0xff, //0x0000019a movzbl       %dil, %edi
	0xc1, 0xe7, 0x08, //0x0000019e shll         $8, %edi
	0x09, 0xcf, //0x000001a1 orl          %ecx, %edi
	0x40, 0x0f, 0xb6, 0xd6, //0x000001a3 movzbl       %sil, %edx
	0x09, 0xfa, //0x000001a7 orl          %edi, %edx
	0x89, 0xd1, //0x000001a9 movl         %edx, %ecx
	0x81, 0xe1, 0xf0, 0xc0, 0xc0, 0x00, //0x000001ab andl         $12632304, %ecx
	0x81, 0xf9, 0xe0, 0x80, 0x80, 0x00, //0x000001b1 cmpl         $8421600, %ecx
	0x0f, 0x85, 0x23, 0x00, 0x00, 0x00, //0x000001b7 jne          LBB0_26
	0x89, 0xd7, //0x000001bd movl         %edx, %edi
	0x81, 0xe7, 0x0f, 0x20, 0x00, 0x00, //0x000001bf andl         $8207, %edi
	0x81, 0xff, 0x0d, 0x20, 0x00, 0x00, //0x000001c5 cmpl         $8205, %edi
	0x0f, 0x84, 0x0f, 0x00, 0x00, 0x00, //0x000001cb je           LBB0_26
	0xb9, 0x03, 0x00, 0x00, 0x00, //0x000001d1 movl         $3, %ecx
	0x85, 0xff, //0x000001d6 testl        %edi, %edi
	0x0f, 0x85, 0x23, 0x00, 0x00, 0x00, //0x000001d8 jne          LBB0_28
	0x90, 0x90, //0x000001de .p2align 4, 0x90
	//0x000001e0 LBB0_26
	0x40, 0xf6, 0xc6, 0x1e, //0x000001e0 testb        $30, %sil
	0x0f, 0x84, 0x07, 0xff, 0xff, 0xff, //0x000001e4 je           LBB0_13
	0x81, 0xe2, 0xe0, 0xc0, 0x00, 0x00, //0x000001ea andl         $49376, %edx
	0xb9, 0x02, 0x00, 0x00, 0x00, //0x000001f0 movl         $2, %ecx
	0x81, 0xfa, 0xc0, 0x80, 0x00, 0x00, //0x000001f5 cmpl         $32960, %edx
	0x0f, 0x85, 0xf0, 0xfe, 0xff, 0xff, //0x000001fb jne          LBB0_13
	//0x00000201 LBB0_28
	0x48, 0x01, 0xc8, //0x00000201 addq         %rcx, %rax
	0x4c, 0x39, 0xd8, //0x00000204 cmpq         %r11, %rax
	0x0f, 0x82, 0x20, 0xff, 0xff, 0xff, //0x00000207 jb           LBB0_17
	//0x0000020d LBB0_30
	0x31, 0xc0, //0x0000020d xorl         %eax, %eax
	0x48, 0x83, 0xc4, 0x08, //0x0000020f addq         $8, %rsp
	0x5b,       //0x00000213 popq         %rbx
	0x5d,       //0x00000214 popq         %rbp
	0xc3,       //0x00000215 retq
	0x00, 0x00, //0x00000216 .p2align 2, 0x00
	//0x00000218 _MASK_USE_NUMBER
	0x02, 0x00, 0x00, 0x00, //0x00000218 .long 2
}
