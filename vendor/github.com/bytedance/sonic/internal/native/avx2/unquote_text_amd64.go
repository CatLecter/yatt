//go:build amd64
// +build amd64

// Code generated by asm2asm, DO NOT EDIT.

package avx2

var _text_unquote = []byte{
	// .p2align 5, 0x00
	// LCPI0_0
	0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, // QUAD $0x5c5c5c5c5c5c5c5c; QUAD $0x5c5c5c5c5c5c5c5c  // .space 16, '\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\'
	0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, //0x00000010 QUAD $0x5c5c5c5c5c5c5c5c; QUAD $0x5c5c5c5c5c5c5c5c  // .space 16, '\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\'
	//0x00000020 .p2align 4, 0x00
	//0x00000020 LCPI0_1
	0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, //0x00000020 QUAD $0x5c5c5c5c5c5c5c5c; QUAD $0x5c5c5c5c5c5c5c5c  // .space 16, '\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\'
	//0x00000030 .p2align 4, 0x90
	//0x00000030 _unquote
	0x55,             //0x00000030 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000031 movq         %rsp, %rbp
	0x41, 0x57, //0x00000034 pushq        %r15
	0x41, 0x56, //0x00000036 pushq        %r14
	0x41, 0x55, //0x00000038 pushq        %r13
	0x41, 0x54, //0x0000003a pushq        %r12
	0x53,                   //0x0000003c pushq        %rbx
	0x48, 0x83, 0xec, 0x20, //0x0000003d subq         $32, %rsp
	0x48, 0x85, 0xf6, //0x00000041 testq        %rsi, %rsi
	0x0f, 0x84, 0xcb, 0x05, 0x00, 0x00, //0x00000044 je           LBB0_1
	0x48, 0x89, 0x4d, 0xd0, //0x0000004a movq         %rcx, $-48(%rbp)
	0x4c, 0x89, 0xc0, //0x0000004e movq         %r8, %rax
	0x4c, 0x89, 0x45, 0xb8, //0x00000051 movq         %r8, $-72(%rbp)
	0x41, 0x83, 0xe0, 0x01, //0x00000055 andl         $1, %r8d
	0x4c, 0x8d, 0x1d, 0x00, 0x08, 0x00, 0x00, //0x00000059 leaq         $2048(%rip), %r11  /* __UnquoteTab+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x0d, 0x98, 0xff, 0xff, 0xff, //0x00000060 vmovdqu      $-104(%rip), %ymm1  /* LCPI0_0+0(%rip) */
	0xc5, 0xfa, 0x6f, 0x15, 0xb0, 0xff, 0xff, 0xff, //0x00000068 vmovdqu      $-80(%rip), %xmm2  /* LCPI0_1+0(%rip) */
	0x48, 0x89, 0x7d, 0xc8, //0x00000070 movq         %rdi, $-56(%rbp)
	0x49, 0x89, 0xf9, //0x00000074 movq         %rdi, %r9
	0x48, 0x89, 0x75, 0xc0, //0x00000077 movq         %rsi, $-64(%rbp)
	0x49, 0x89, 0xf2, //0x0000007b movq         %rsi, %r10
	0x4c, 0x89, 0xc6, //0x0000007e movq         %r8, %rsi
	0x49, 0x89, 0xd0, //0x00000081 movq         %rdx, %r8
	0xe9, 0x77, 0x02, 0x00, 0x00, //0x00000084 jmp          LBB0_3
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000089 .p2align 4, 0x90
	//0x00000090 LBB0_60
	0x41, 0x81, 0xff, 0xff, 0x07, 0x00, 0x00, //0x00000090 cmpl         $2047, %r15d
	0x0f, 0x86, 0xbd, 0x01, 0x00, 0x00, //0x00000097 jbe          LBB0_61
	0x41, 0x8d, 0x87, 0x00, 0x20, 0xff, 0xff, //0x0000009d leal         $-57344(%r15), %eax
	0x3d, 0xff, 0xf7, 0xff, 0xff, //0x000000a4 cmpl         $-2049, %eax
	0x0f, 0x86, 0x71, 0x01, 0x00, 0x00, //0x000000a9 jbe          LBB0_63
	0x48, 0x85, 0xf6, //0x000000af testq        %rsi, %rsi
	0x0f, 0x85, 0x14, 0x01, 0x00, 0x00, //0x000000b2 jne          LBB0_65
	0x49, 0x83, 0xfa, 0x06, //0x000000b8 cmpq         $6, %r10
	0x0f, 0x8c, 0x2f, 0x01, 0x00, 0x00, //0x000000bc jl           LBB0_74
	//0x000000c2 LBB0_71
	0x41, 0x81, 0xff, 0xff, 0xdb, 0x00, 0x00, //0x000000c2 cmpl         $56319, %r15d
	0x0f, 0x87, 0x22, 0x01, 0x00, 0x00, //0x000000c9 ja           LBB0_74
	0x41, 0x80, 0x39, 0x5c, //0x000000cf cmpb         $92, (%r9)
	0x0f, 0x85, 0x18, 0x01, 0x00, 0x00, //0x000000d3 jne          LBB0_74
	0x41, 0x80, 0x79, 0x01, 0x75, //0x000000d9 cmpb         $117, $1(%r9)
	0x0f, 0x85, 0x0d, 0x01, 0x00, 0x00, //0x000000de jne          LBB0_74
	0x45, 0x8b, 0x71, 0x02, //0x000000e4 movl         $2(%r9), %r14d
	0x44, 0x89, 0xf1, //0x000000e8 movl         %r14d, %ecx
	0xf7, 0xd1, //0x000000eb notl         %ecx
	0x41, 0x8d, 0x86, 0xd0, 0xcf, 0xcf, 0xcf, //0x000000ed leal         $-808464432(%r14), %eax
	0x81, 0xe1, 0x80, 0x80, 0x80, 0x80, //0x000000f4 andl         $-2139062144, %ecx
	0x85, 0xc1, //0x000000fa testl        %eax, %ecx
	0x0f, 0x85, 0x02, 0x06, 0x00, 0x00, //0x000000fc jne          LBB0_84
	0x41, 0x8d, 0x86, 0x19, 0x19, 0x19, 0x19, //0x00000102 leal         $421075225(%r14), %eax
	0x44, 0x09, 0xf0, //0x00000109 orl          %r14d, %eax
	0xa9, 0x80, 0x80, 0x80, 0x80, //0x0000010c testl        $-2139062144, %eax
	0x0f, 0x85, 0xed, 0x05, 0x00, 0x00, //0x00000111 jne          LBB0_84
	0x44, 0x89, 0xf3, //0x00000117 movl         %r14d, %ebx
	0x81, 0xe3, 0x7f, 0x7f, 0x7f, 0x7f, //0x0000011a andl         $2139062143, %ebx
	0xb8, 0xc0, 0xc0, 0xc0, 0xc0, //0x00000120 movl         $-1061109568, %eax
	0x29, 0xd8, //0x00000125 subl         %ebx, %eax
	0x8d, 0x93, 0x46, 0x46, 0x46, 0x46, //0x00000127 leal         $1179010630(%rbx), %edx
	0x21, 0xc8, //0x0000012d andl         %ecx, %eax
	0x85, 0xd0, //0x0000012f testl        %edx, %eax
	0x0f, 0x85, 0xcd, 0x05, 0x00, 0x00, //0x00000131 jne          LBB0_84
	0xb8, 0xe0, 0xe0, 0xe0, 0xe0, //0x00000137 movl         $-522133280, %eax
	0x29, 0xd8, //0x0000013c subl         %ebx, %eax
	0x81, 0xc3, 0x39, 0x39, 0x39, 0x39, //0x0000013e addl         $960051513, %ebx
	0x21, 0xc1, //0x00000144 andl         %eax, %ecx
	0x85, 0xd9, //0x00000146 testl        %ebx, %ecx
	0x0f, 0x85, 0xb6, 0x05, 0x00, 0x00, //0x00000148 jne          LBB0_84
	0x41, 0x0f, 0xce, //0x0000014e bswapl       %r14d
	0x44, 0x89, 0xf0, //0x00000151 movl         %r14d, %eax
	0xc1, 0xe8, 0x04, //0x00000154 shrl         $4, %eax
	0xf7, 0xd0, //0x00000157 notl         %eax
	0x25, 0x01, 0x01, 0x01, 0x01, //0x00000159 andl         $16843009, %eax
	0x8d, 0x04, 0xc0, //0x0000015e leal         (%rax,%rax,8), %eax
	0x41, 0x81, 0xe6, 0x0f, 0x0f, 0x0f, 0x0f, //0x00000161 andl         $252645135, %r14d
	0x41, 0x01, 0xc6, //0x00000168 addl         %eax, %r14d
	0x44, 0x89, 0xf0, //0x0000016b movl         %r14d, %eax
	0xc1, 0xe8, 0x04, //0x0000016e shrl         $4, %eax
	0x44, 0x09, 0xf0, //0x00000171 orl          %r14d, %eax
	0x0f, 0xb6, 0xc8, //0x00000174 movzbl       %al, %ecx
	0xc1, 0xe8, 0x08, //0x00000177 shrl         $8, %eax
	0x25, 0x00, 0xff, 0x00, 0x00, //0x0000017a andl         $65280, %eax
	0x8d, 0x1c, 0x08, //0x0000017f leal         (%rax,%rcx), %ebx
	0x49, 0x83, 0xc1, 0x06, //0x00000182 addq         $6, %r9
	0x49, 0x83, 0xc2, 0xfa, //0x00000186 addq         $-6, %r10
	0x01, 0xc8, //0x0000018a addl         %ecx, %eax
	0x05, 0x00, 0x20, 0xff, 0xff, //0x0000018c addl         $-57344, %eax
	0x3d, 0xff, 0xfb, 0xff, 0xff, //0x00000191 cmpl         $-1025, %eax
	0x0f, 0x87, 0xe7, 0x00, 0x00, 0x00, //0x00000196 ja           LBB0_96
	0xf6, 0x45, 0xb8, 0x02, //0x0000019c testb        $2, $-72(%rbp)
	0x0f, 0x84, 0x4e, 0x06, 0x00, 0x00, //0x000001a0 je           LBB0_78
	0x4c, 0x89, 0xda, //0x000001a6 movq         %r11, %rdx
	0x66, 0x41, 0xc7, 0x40, 0xfe, 0xef, 0xbf, //0x000001a9 movw         $-16401, $-2(%r8)
	0x41, 0xc6, 0x00, 0xbd, //0x000001b0 movb         $-67, (%r8)
	0x49, 0x83, 0xc0, 0x03, //0x000001b4 addq         $3, %r8
	0x41, 0x89, 0xdf, //0x000001b8 movl         %ebx, %r15d
	0x81, 0xfb, 0x80, 0x00, 0x00, 0x00, //0x000001bb cmpl         $128, %ebx
	0x0f, 0x83, 0xc9, 0xfe, 0xff, 0xff, //0x000001c1 jae          LBB0_60
	0xe9, 0x10, 0x01, 0x00, 0x00, //0x000001c7 jmp          LBB0_57
	//0x000001cc LBB0_65
	0x4d, 0x85, 0xd2, //0x000001cc testq        %r10, %r10
	0x0f, 0x8e, 0x46, 0x06, 0x00, 0x00, //0x000001cf jle          LBB0_66
	0x41, 0x80, 0x39, 0x5c, //0x000001d5 cmpb         $92, (%r9)
	0x0f, 0x85, 0x09, 0x01, 0x00, 0x00, //0x000001d9 jne          LBB0_76
	0x49, 0x83, 0xc2, 0xff, //0x000001df addq         $-1, %r10
	0x49, 0x83, 0xc1, 0x01, //0x000001e3 addq         $1, %r9
	0x49, 0x83, 0xfa, 0x06, //0x000001e7 cmpq         $6, %r10
	0x0f, 0x8d, 0xd1, 0xfe, 0xff, 0xff, //0x000001eb jge          LBB0_71
	//0x000001f1 LBB0_74
	0xf6, 0x45, 0xb8, 0x02, //0x000001f1 testb        $2, $-72(%rbp)
	0x0f, 0x84, 0x14, 0x06, 0x00, 0x00, //0x000001f5 je           LBB0_75
	//0x000001fb LBB0_77
	0x66, 0x41, 0xc7, 0x40, 0xfe, 0xef, 0xbf, //0x000001fb movw         $-16401, $-2(%r8)
	0x41, 0xc6, 0x00, 0xbd, //0x00000202 movb         $-67, (%r8)
	0x49, 0x83, 0xc0, 0x01, //0x00000206 addq         $1, %r8
	0x4c, 0x89, 0xda, //0x0000020a movq         %r11, %rdx
	0x49, 0x89, 0xfb, //0x0000020d movq         %rdi, %r11
	0x4d, 0x85, 0xd2, //0x00000210 testq        %r10, %r10
	0x0f, 0x85, 0xe7, 0x00, 0x00, 0x00, //0x00000213 jne          LBB0_3
	0xe9, 0xef, 0x03, 0x00, 0x00, //0x00000219 jmp          LBB0_99
	0x90, 0x90, //0x0000021e .p2align 4, 0x90
	//0x00000220 LBB0_63
	0x44, 0x89, 0xf8, //0x00000220 movl         %r15d, %eax
	0xc1, 0xe8, 0x0c, //0x00000223 shrl         $12, %eax
	0x0c, 0xe0, //0x00000226 orb          $-32, %al
	0x41, 0x88, 0x40, 0xfe, //0x00000228 movb         %al, $-2(%r8)
	0x44, 0x89, 0xf8, //0x0000022c movl         %r15d, %eax
	0xc1, 0xe8, 0x06, //0x0000022f shrl         $6, %eax
	0x24, 0x3f, //0x00000232 andb         $63, %al
	0x0c, 0x80, //0x00000234 orb          $-128, %al
	0x41, 0x88, 0x40, 0xff, //0x00000236 movb         %al, $-1(%r8)
	0x41, 0x80, 0xe7, 0x3f, //0x0000023a andb         $63, %r15b
	0x41, 0x80, 0xcf, 0x80, //0x0000023e orb          $-128, %r15b
	0x45, 0x88, 0x38, //0x00000242 movb         %r15b, (%r8)
	0x49, 0x83, 0xc0, 0x01, //0x00000245 addq         $1, %r8
	0x49, 0x89, 0xfb, //0x00000249 movq         %rdi, %r11
	0x4d, 0x85, 0xd2, //0x0000024c testq        %r10, %r10
	0x0f, 0x85, 0xab, 0x00, 0x00, 0x00, //0x0000024f jne          LBB0_3
	0xe9, 0xb3, 0x03, 0x00, 0x00, //0x00000255 jmp          LBB0_99
	//0x0000025a LBB0_61
	0x44, 0x89, 0xf8, //0x0000025a movl         %r15d, %eax
	0xc1, 0xe8, 0x06, //0x0000025d shrl         $6, %eax
	0x0c, 0xc0, //0x00000260 orb          $-64, %al
	0x41, 0x88, 0x40, 0xfe, //0x00000262 movb         %al, $-2(%r8)
	0x41, 0x80, 0xe7, 0x3f, //0x00000266 andb         $63, %r15b
	0x41, 0x80, 0xcf, 0x80, //0x0000026a orb          $-128, %r15b
	0x45, 0x88, 0x78, 0xff, //0x0000026e movb         %r15b, $-1(%r8)
	0x49, 0x89, 0xfb, //0x00000272 movq         %rdi, %r11
	0x4d, 0x85, 0xd2, //0x00000275 testq        %r10, %r10
	0x0f, 0x85, 0x82, 0x00, 0x00, 0x00, //0x00000278 jne          LBB0_3
	0xe9, 0x8a, 0x03, 0x00, 0x00, //0x0000027e jmp          LBB0_99
	//0x00000283 LBB0_96
	0x41, 0xc1, 0xe7, 0x0a, //0x00000283 shll         $10, %r15d
	0x89, 0xd9, //0x00000287 movl         %ebx, %ecx
	0x44, 0x01, 0xf9, //0x00000289 addl         %r15d, %ecx
	0x42, 0x8d, 0x04, 0x3b, //0x0000028c leal         (%rbx,%r15), %eax
	0x05, 0x00, 0x24, 0xa0, 0xfc, //0x00000290 addl         $-56613888, %eax
	0x89, 0xc2, //0x00000295 movl         %eax, %edx
	0xc1, 0xea, 0x12, //0x00000297 shrl         $18, %edx
	0x80, 0xca, 0xf0, //0x0000029a orb          $-16, %dl
	0x41, 0x88, 0x50, 0xfe, //0x0000029d movb         %dl, $-2(%r8)
	0x89, 0xc2, //0x000002a1 movl         %eax, %edx
	0xc1, 0xea, 0x0c, //0x000002a3 shrl         $12, %edx
	0x80, 0xe2, 0x3f, //0x000002a6 andb         $63, %dl
	0x80, 0xca, 0x80, //0x000002a9 orb          $-128, %dl
	0x41, 0x88, 0x50, 0xff, //0x000002ac movb         %dl, $-1(%r8)
	0x4c, 0x89, 0xda, //0x000002b0 movq         %r11, %rdx
	0xc1, 0xe8, 0x06, //0x000002b3 shrl         $6, %eax
	0x24, 0x3f, //0x000002b6 andb         $63, %al
	0x0c, 0x80, //0x000002b8 orb          $-128, %al
	0x41, 0x88, 0x00, //0x000002ba movb         %al, (%r8)
	0x80, 0xe1, 0x3f, //0x000002bd andb         $63, %cl
	0x80, 0xc9, 0x80, //0x000002c0 orb          $-128, %cl
	0x41, 0x88, 0x48, 0x01, //0x000002c3 movb         %cl, $1(%r8)
	0x49, 0x83, 0xc0, 0x02, //0x000002c7 addq         $2, %r8
	0x49, 0x89, 0xfb, //0x000002cb movq         %rdi, %r11
	0x4d, 0x85, 0xd2, //0x000002ce testq        %r10, %r10
	0x0f, 0x85, 0x29, 0x00, 0x00, 0x00, //0x000002d1 jne          LBB0_3
	0xe9, 0x31, 0x03, 0x00, 0x00, //0x000002d7 jmp          LBB0_99
	//0x000002dc LBB0_57
	0x49, 0x83, 0xc0, 0xfe, //0x000002dc addq         $-2, %r8
	0x4d, 0x89, 0xc4, //0x000002e0 movq         %r8, %r12
	0xe9, 0x0e, 0x03, 0x00, 0x00, //0x000002e3 jmp          LBB0_58
	//0x000002e8 LBB0_76
	0xf6, 0x45, 0xb8, 0x02, //0x000002e8 testb        $2, $-72(%rbp)
	0x0f, 0x85, 0x09, 0xff, 0xff, 0xff, //0x000002ec jne          LBB0_77
	0xe9, 0xfd, 0x04, 0x00, 0x00, //0x000002f2 jmp          LBB0_78
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000002f7 .p2align 4, 0x90
	//0x00000300 LBB0_3
	0x41, 0x80, 0x39, 0x5c, //0x00000300 cmpb         $92, (%r9)
	0x0f, 0x85, 0x16, 0x00, 0x00, 0x00, //0x00000304 jne          LBB0_5
	0x45, 0x31, 0xed, //0x0000030a xorl         %r13d, %r13d
	0xe9, 0x5e, 0x01, 0x00, 0x00, //0x0000030d jmp          LBB0_23
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000312 .p2align 4, 0x90
	//0x00000320 LBB0_5
	0x4d, 0x89, 0xd7, //0x00000320 movq         %r10, %r15
	0x4c, 0x89, 0xc0, //0x00000323 movq         %r8, %rax
	0x4d, 0x89, 0xcd, //0x00000326 movq         %r9, %r13
	0x49, 0x83, 0xfa, 0x20, //0x00000329 cmpq         $32, %r10
	0x0f, 0x8c, 0x45, 0x00, 0x00, 0x00, //0x0000032d jl           LBB0_11
	0x31, 0xc0, //0x00000333 xorl         %eax, %eax
	0x4c, 0x89, 0xd1, //0x00000335 movq         %r10, %rcx
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000338 .p2align 4, 0x90
	//0x00000340 LBB0_7
	0xc4, 0xc1, 0x7e, 0x6f, 0x04, 0x01, //0x00000340 vmovdqu      (%r9,%rax), %ymm0
	0xc4, 0xc1, 0x7e, 0x7f, 0x04, 0x00, //0x00000346 vmovdqu      %ymm0, (%r8,%rax)
	0xc5, 0xfd, 0x74, 0xc1, //0x0000034c vpcmpeqb     %ymm1, %ymm0, %ymm0
	0xc5, 0xfd, 0xd7, 0xd8, //0x00000350 vpmovmskb    %ymm0, %ebx
	0x85, 0xdb, //0x00000354 testl        %ebx, %ebx
	0x0f, 0x85, 0xdd, 0x00, 0x00, 0x00, //0x00000356 jne          LBB0_8
	0x4c, 0x8d, 0x79, 0xe0, //0x0000035c leaq         $-32(%rcx), %r15
	0x48, 0x83, 0xc0, 0x20, //0x00000360 addq         $32, %rax
	0x48, 0x83, 0xf9, 0x3f, //0x00000364 cmpq         $63, %rcx
	0x4c, 0x89, 0xf9, //0x00000368 movq         %r15, %rcx
	0x0f, 0x8f, 0xcf, 0xff, 0xff, 0xff, //0x0000036b jg           LBB0_7
	0x4d, 0x8d, 0x2c, 0x01, //0x00000371 leaq         (%r9,%rax), %r13
	0x4c, 0x01, 0xc0, //0x00000375 addq         %r8, %rax
	//0x00000378 LBB0_11
	0xc5, 0xf8, 0x77, //0x00000378 vzeroupper
	0x49, 0x83, 0xff, 0x10, //0x0000037b cmpq         $16, %r15
	0x0f, 0x8c, 0x57, 0x00, 0x00, 0x00, //0x0000037f jl           LBB0_12
	0x4d, 0x89, 0xcc, //0x00000385 movq         %r9, %r12
	0x4d, 0x29, 0xec, //0x00000388 subq         %r13, %r12
	0xc5, 0xfe, 0x6f, 0x0d, 0x6d, 0xfc, 0xff, 0xff, //0x0000038b vmovdqu      $-915(%rip), %ymm1  /* LCPI0_0+0(%rip) */
	0xc5, 0xfa, 0x6f, 0x15, 0x85, 0xfc, 0xff, 0xff, //0x00000393 vmovdqu      $-891(%rip), %xmm2  /* LCPI0_1+0(%rip) */
	0x90, 0x90, 0x90, 0x90, 0x90, //0x0000039b .p2align 4, 0x90
	//0x000003a0 LBB0_18
	0xc4, 0xc1, 0x7a, 0x6f, 0x45, 0x00, //0x000003a0 vmovdqu      (%r13), %xmm0
	0xc5, 0xfa, 0x7f, 0x00, //0x000003a6 vmovdqu      %xmm0, (%rax)
	0xc5, 0xf9, 0x74, 0xc2, //0x000003aa vpcmpeqb     %xmm2, %xmm0, %xmm0
	0xc5, 0xf9, 0xd7, 0xc8, //0x000003ae vpmovmskb    %xmm0, %ecx
	0x85, 0xc9, //0x000003b2 testl        %ecx, %ecx
	0x0f, 0x85, 0x95, 0x00, 0x00, 0x00, //0x000003b4 jne          LBB0_19
	0x49, 0x83, 0xc5, 0x10, //0x000003ba addq         $16, %r13
	0x48, 0x83, 0xc0, 0x10, //0x000003be addq         $16, %rax
	0x4d, 0x8d, 0x77, 0xf0, //0x000003c2 leaq         $-16(%r15), %r14
	0x49, 0x83, 0xc4, 0xf0, //0x000003c6 addq         $-16, %r12
	0x49, 0x83, 0xff, 0x1f, //0x000003ca cmpq         $31, %r15
	0x4d, 0x89, 0xf7, //0x000003ce movq         %r14, %r15
	0x0f, 0x87, 0xc9, 0xff, 0xff, 0xff, //0x000003d1 ja           LBB0_18
	0xe9, 0x13, 0x00, 0x00, 0x00, //0x000003d7 jmp          LBB0_13
	//0x000003dc LBB0_12
	0x4d, 0x89, 0xfe, //0x000003dc movq         %r15, %r14
	0xc5, 0xfe, 0x6f, 0x0d, 0x19, 0xfc, 0xff, 0xff, //0x000003df vmovdqu      $-999(%rip), %ymm1  /* LCPI0_0+0(%rip) */
	0xc5, 0xfa, 0x6f, 0x15, 0x31, 0xfc, 0xff, 0xff, //0x000003e7 vmovdqu      $-975(%rip), %xmm2  /* LCPI0_1+0(%rip) */
	//0x000003ef LBB0_13
	0x4d, 0x85, 0xf6, //0x000003ef testq        %r14, %r14
	0x0f, 0x84, 0x23, 0x02, 0x00, 0x00, //0x000003f2 je           LBB0_100
	0x31, 0xc9, //0x000003f8 xorl         %ecx, %ecx
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000003fa .p2align 4, 0x90
	//0x00000400 LBB0_15
	0x41, 0x0f, 0xb6, 0x5c, 0x0d, 0x00, //0x00000400 movzbl       (%r13,%rcx), %ebx
	0x80, 0xfb, 0x5c, //0x00000406 cmpb         $92, %bl
	0x0f, 0x84, 0x15, 0x00, 0x00, 0x00, //0x00000409 je           LBB0_21
	0x88, 0x1c, 0x08, //0x0000040f movb         %bl, (%rax,%rcx)
	0x48, 0x83, 0xc1, 0x01, //0x00000412 addq         $1, %rcx
	0x49, 0x39, 0xce, //0x00000416 cmpq         %rcx, %r14
	0x0f, 0x85, 0xe1, 0xff, 0xff, 0xff, //0x00000419 jne          LBB0_15
	0xe9, 0xf7, 0x01, 0x00, 0x00, //0x0000041f jmp          LBB0_100
	//0x00000424 LBB0_21
	0x4d, 0x29, 0xcd, //0x00000424 subq         %r9, %r13
	0x49, 0x01, 0xcd, //0x00000427 addq         %rcx, %r13
	0x49, 0x83, 0xfd, 0xff, //0x0000042a cmpq         $-1, %r13
	0x0f, 0x85, 0x3c, 0x00, 0x00, 0x00, //0x0000042e jne          LBB0_23
	0xe9, 0xe2, 0x01, 0x00, 0x00, //0x00000434 jmp          LBB0_100
	//0x00000439 LBB0_8
	0x44, 0x0f, 0xbc, 0xeb, //0x00000439 bsfl         %ebx, %r13d
	0x49, 0x01, 0xc5, //0x0000043d addq         %rax, %r13
	0x49, 0x83, 0xfd, 0xff, //0x00000440 cmpq         $-1, %r13
	0x0f, 0x85, 0x26, 0x00, 0x00, 0x00, //0x00000444 jne          LBB0_23
	0xe9, 0xcc, 0x01, 0x00, 0x00, //0x0000044a jmp          LBB0_100
	//0x0000044f LBB0_19
	0x66, 0x0f, 0xbc, 0xc1, //0x0000044f bsfw         %cx, %ax
	0x44, 0x0f, 0xb7, 0xe8, //0x00000453 movzwl       %ax, %r13d
	0x4d, 0x29, 0xe5, //0x00000457 subq         %r12, %r13
	0x49, 0x83, 0xfd, 0xff, //0x0000045a cmpq         $-1, %r13
	0x0f, 0x84, 0xb7, 0x01, 0x00, 0x00, //0x0000045e je           LBB0_100
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000464 .p2align 4, 0x90
	//0x00000470 LBB0_23
	0x49, 0x8d, 0x45, 0x02, //0x00000470 leaq         $2(%r13), %rax
	0x49, 0x29, 0xc2, //0x00000474 subq         %rax, %r10
	0x0f, 0x88, 0xc2, 0x03, 0x00, 0x00, //0x00000477 js           LBB0_24
	0x4d, 0x01, 0xe9, //0x0000047d addq         %r13, %r9
	0x49, 0x83, 0xc1, 0x02, //0x00000480 addq         $2, %r9
	0x48, 0x85, 0xf6, //0x00000484 testq        %rsi, %rsi
	0x0f, 0x85, 0x15, 0x01, 0x00, 0x00, //0x00000487 jne          LBB0_26
	//0x0000048d LBB0_37
	0x4f, 0x8d, 0x24, 0x28, //0x0000048d leaq         (%r8,%r13), %r12
	0x41, 0x0f, 0xb6, 0x41, 0xff, //0x00000491 movzbl       $-1(%r9), %eax
	0x42, 0x8a, 0x0c, 0x18, //0x00000496 movb         (%rax,%r11), %cl
	0x80, 0xf9, 0xff, //0x0000049a cmpb         $-1, %cl
	0x0f, 0x84, 0x2d, 0x00, 0x00, 0x00, //0x0000049d je           LBB0_41
	0x84, 0xc9, //0x000004a3 testb        %cl, %cl
	0x0f, 0x84, 0x3c, 0x02, 0x00, 0x00, //0x000004a5 je           LBB0_39
	0x41, 0x88, 0x0c, 0x24, //0x000004ab movb         %cl, (%r12)
	0x49, 0x83, 0xc4, 0x01, //0x000004af addq         $1, %r12
	0x4d, 0x89, 0xe0, //0x000004b3 movq         %r12, %r8
	0x4d, 0x85, 0xd2, //0x000004b6 testq        %r10, %r10
	0x0f, 0x85, 0x41, 0xfe, 0xff, 0xff, //0x000004b9 jne          LBB0_3
	0xe9, 0x49, 0x01, 0x00, 0x00, //0x000004bf jmp          LBB0_99
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000004c4 .p2align 4, 0x90
	//0x000004d0 LBB0_41
	0x49, 0x83, 0xfa, 0x03, //0x000004d0 cmpq         $3, %r10
	0x0f, 0x86, 0x65, 0x03, 0x00, 0x00, //0x000004d4 jbe          LBB0_24
	0x45, 0x8b, 0x31, //0x000004da movl         (%r9), %r14d
	0x44, 0x89, 0xf1, //0x000004dd movl         %r14d, %ecx
	0xf7, 0xd1, //0x000004e0 notl         %ecx
	0x41, 0x8d, 0x86, 0xd0, 0xcf, 0xcf, 0xcf, //0x000004e2 leal         $-808464432(%r14), %eax
	0x81, 0xe1, 0x80, 0x80, 0x80, 0x80, //0x000004e9 andl         $-2139062144, %ecx
	0x85, 0xc1, //0x000004ef testl        %eax, %ecx
	0x0f, 0x85, 0x3f, 0x01, 0x00, 0x00, //0x000004f1 jne          LBB0_46
	0x41, 0x8d, 0x86, 0x19, 0x19, 0x19, 0x19, //0x000004f7 leal         $421075225(%r14), %eax
	0x44, 0x09, 0xf0, //0x000004fe orl          %r14d, %eax
	0xa9, 0x80, 0x80, 0x80, 0x80, //0x00000501 testl        $-2139062144, %eax
	0x0f, 0x85, 0x2a, 0x01, 0x00, 0x00, //0x00000506 jne          LBB0_46
	0x4c, 0x89, 0xdf, //0x0000050c movq         %r11, %rdi
	0x44, 0x89, 0xf3, //0x0000050f movl         %r14d, %ebx
	0x81, 0xe3, 0x7f, 0x7f, 0x7f, 0x7f, //0x00000512 andl         $2139062143, %ebx
	0xb8, 0xc0, 0xc0, 0xc0, 0xc0, //0x00000518 movl         $-1061109568, %eax
	0x29, 0xd8, //0x0000051d subl         %ebx, %eax
	0x49, 0x89, 0xd3, //0x0000051f movq         %rdx, %r11
	0x8d, 0x93, 0x46, 0x46, 0x46, 0x46, //0x00000522 leal         $1179010630(%rbx), %edx
	0x21, 0xc8, //0x00000528 andl         %ecx, %eax
	0x85, 0xd0, //0x0000052a testl        %edx, %eax
	0x0f, 0x85, 0x04, 0x01, 0x00, 0x00, //0x0000052c jne          LBB0_46
	0xb8, 0xe0, 0xe0, 0xe0, 0xe0, //0x00000532 movl         $-522133280, %eax
	0x29, 0xd8, //0x00000537 subl         %ebx, %eax
	0x81, 0xc3, 0x39, 0x39, 0x39, 0x39, //0x00000539 addl         $960051513, %ebx
	0x21, 0xc1, //0x0000053f andl         %eax, %ecx
	0x85, 0xd9, //0x00000541 testl        %ebx, %ecx
	0x0f, 0x85, 0xed, 0x00, 0x00, 0x00, //0x00000543 jne          LBB0_46
	0x4c, 0x89, 0xda, //0x00000549 movq         %r11, %rdx
	0x41, 0x0f, 0xce, //0x0000054c bswapl       %r14d
	0x44, 0x89, 0xf0, //0x0000054f movl         %r14d, %eax
	0xc1, 0xe8, 0x04, //0x00000552 shrl         $4, %eax
	0xf7, 0xd0, //0x00000555 notl         %eax
	0x25, 0x01, 0x01, 0x01, 0x01, //0x00000557 andl         $16843009, %eax
	0x8d, 0x04, 0xc0, //0x0000055c leal         (%rax,%rax,8), %eax
	0x41, 0x81, 0xe6, 0x0f, 0x0f, 0x0f, 0x0f, //0x0000055f andl         $252645135, %r14d
	0x41, 0x01, 0xc6, //0x00000566 addl         %eax, %r14d
	0x44, 0x89, 0xf0, //0x00000569 movl         %r14d, %eax
	0xc1, 0xe8, 0x04, //0x0000056c shrl         $4, %eax
	0x44, 0x09, 0xf0, //0x0000056f orl          %r14d, %eax
	0x44, 0x0f, 0xb6, 0xf8, //0x00000572 movzbl       %al, %r15d
	0xc1, 0xe8, 0x08, //0x00000576 shrl         $8, %eax
	0x25, 0x00, 0xff, 0x00, 0x00, //0x00000579 andl         $65280, %eax
	0x41, 0x09, 0xc7, //0x0000057e orl          %eax, %r15d
	0x49, 0x83, 0xc1, 0x04, //0x00000581 addq         $4, %r9
	0x49, 0x83, 0xc2, 0xfc, //0x00000585 addq         $-4, %r10
	0x41, 0x81, 0xff, 0x80, 0x00, 0x00, 0x00, //0x00000589 cmpl         $128, %r15d
	0x0f, 0x82, 0x5d, 0x00, 0x00, 0x00, //0x00000590 jb           LBB0_56
	0x4d, 0x01, 0xe8, //0x00000596 addq         %r13, %r8
	0x49, 0x83, 0xc0, 0x02, //0x00000599 addq         $2, %r8
	0xe9, 0xee, 0xfa, 0xff, 0xff, //0x0000059d jmp          LBB0_60
	//0x000005a2 LBB0_26
	0x45, 0x85, 0xd2, //0x000005a2 testl        %r10d, %r10d
	0x0f, 0x84, 0x94, 0x02, 0x00, 0x00, //0x000005a5 je           LBB0_24
	0x41, 0x80, 0x79, 0xff, 0x5c, //0x000005ab cmpb         $92, $-1(%r9)
	0x0f, 0x85, 0x14, 0x02, 0x00, 0x00, //0x000005b0 jne          LBB0_28
	0x41, 0x80, 0x39, 0x5c, //0x000005b6 cmpb         $92, (%r9)
	0x0f, 0x85, 0x26, 0x00, 0x00, 0x00, //0x000005ba jne          LBB0_36
	0x41, 0x83, 0xfa, 0x01, //0x000005c0 cmpl         $1, %r10d
	0x0f, 0x8e, 0x75, 0x02, 0x00, 0x00, //0x000005c4 jle          LBB0_24
	0x41, 0x8a, 0x41, 0x01, //0x000005ca movb         $1(%r9), %al
	0x3c, 0x22, //0x000005ce cmpb         $34, %al
	0x0f, 0x84, 0x08, 0x00, 0x00, 0x00, //0x000005d0 je           LBB0_35
	0x3c, 0x5c, //0x000005d6 cmpb         $92, %al
	0x0f, 0x85, 0x09, 0x02, 0x00, 0x00, //0x000005d8 jne          LBB0_33
	//0x000005de LBB0_35
	0x49, 0x83, 0xc1, 0x01, //0x000005de addq         $1, %r9
	0x49, 0x83, 0xc2, 0xff, //0x000005e2 addq         $-1, %r10
	//0x000005e6 LBB0_36
	0x49, 0x83, 0xc1, 0x01, //0x000005e6 addq         $1, %r9
	0x49, 0x83, 0xc2, 0xff, //0x000005ea addq         $-1, %r10
	0xe9, 0x9a, 0xfe, 0xff, 0xff, //0x000005ee jmp          LBB0_37
	//0x000005f3 LBB0_56
	0x44, 0x89, 0xfb, //0x000005f3 movl         %r15d, %ebx
	//0x000005f6 LBB0_58
	0x41, 0x88, 0x1c, 0x24, //0x000005f6 movb         %bl, (%r12)
	0x49, 0x83, 0xc4, 0x01, //0x000005fa addq         $1, %r12
	0x4d, 0x89, 0xe0, //0x000005fe movq         %r12, %r8
	0x49, 0x89, 0xfb, //0x00000601 movq         %rdi, %r11
	0x4d, 0x85, 0xd2, //0x00000604 testq        %r10, %r10
	0x0f, 0x85, 0xf3, 0xfc, 0xff, 0xff, //0x00000607 jne          LBB0_3
	//0x0000060d LBB0_99
	0x45, 0x31, 0xd2, //0x0000060d xorl         %r10d, %r10d
	0xe9, 0x06, 0x00, 0x00, 0x00, //0x00000610 jmp          LBB0_100
	//0x00000615 LBB0_1
	0x45, 0x31, 0xd2, //0x00000615 xorl         %r10d, %r10d
	0x49, 0x89, 0xd0, //0x00000618 movq         %rdx, %r8
	//0x0000061b LBB0_100
	0x4d, 0x01, 0xd0, //0x0000061b addq         %r10, %r8
	0x49, 0x29, 0xd0, //0x0000061e subq         %rdx, %r8
	//0x00000621 LBB0_101
	0x4c, 0x89, 0xc0, //0x00000621 movq         %r8, %rax
	0x48, 0x83, 0xc4, 0x20, //0x00000624 addq         $32, %rsp
	0x5b,       //0x00000628 popq         %rbx
	0x41, 0x5c, //0x00000629 popq         %r12
	0x41, 0x5d, //0x0000062b popq         %r13
	0x41, 0x5e, //0x0000062d popq         %r14
	0x41, 0x5f, //0x0000062f popq         %r15
	0x5d,             //0x00000631 popq         %rbp
	0xc5, 0xf8, 0x77, //0x00000632 vzeroupper
	0xc3, //0x00000635 retq
	//0x00000636 LBB0_46
	0x4c, 0x89, 0xc8, //0x00000636 movq         %r9, %rax
	0x48, 0x2b, 0x45, 0xc8, //0x00000639 subq         $-56(%rbp), %rax
	0x48, 0x8b, 0x75, 0xd0, //0x0000063d movq         $-48(%rbp), %rsi
	0x48, 0x89, 0x06, //0x00000641 movq         %rax, (%rsi)
	0x41, 0x8a, 0x09, //0x00000644 movb         (%r9), %cl
	0x8d, 0x51, 0xc6, //0x00000647 leal         $-58(%rcx), %edx
	0x49, 0xc7, 0xc0, 0xfe, 0xff, 0xff, 0xff, //0x0000064a movq         $-2, %r8
	0x80, 0xfa, 0xf5, //0x00000651 cmpb         $-11, %dl
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x00000654 ja           LBB0_48
	0x80, 0xe1, 0xdf, //0x0000065a andb         $-33, %cl
	0x80, 0xc1, 0xb9, //0x0000065d addb         $-71, %cl
	0x80, 0xf9, 0xfa, //0x00000660 cmpb         $-6, %cl
	0x0f, 0x82, 0xb8, 0xff, 0xff, 0xff, //0x00000663 jb           LBB0_101
	//0x00000669 LBB0_48
	0x48, 0x8d, 0x48, 0x01, //0x00000669 leaq         $1(%rax), %rcx
	0x48, 0x89, 0x0e, //0x0000066d movq         %rcx, (%rsi)
	0x41, 0x8a, 0x49, 0x01, //0x00000670 movb         $1(%r9), %cl
	0x8d, 0x51, 0xc6, //0x00000674 leal         $-58(%rcx), %edx
	0x80, 0xfa, 0xf5, //0x00000677 cmpb         $-11, %dl
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x0000067a ja           LBB0_50
	0x80, 0xe1, 0xdf, //0x00000680 andb         $-33, %cl
	0x80, 0xc1, 0xb9, //0x00000683 addb         $-71, %cl
	0x80, 0xf9, 0xfa, //0x00000686 cmpb         $-6, %cl
	0x0f, 0x82, 0x92, 0xff, 0xff, 0xff, //0x00000689 jb           LBB0_101
	//0x0000068f LBB0_50
	0x48, 0x8d, 0x48, 0x02, //0x0000068f leaq         $2(%rax), %rcx
	0x48, 0x89, 0x0e, //0x00000693 movq         %rcx, (%rsi)
	0x41, 0x8a, 0x49, 0x02, //0x00000696 movb         $2(%r9), %cl
	0x8d, 0x51, 0xc6, //0x0000069a leal         $-58(%rcx), %edx
	0x80, 0xfa, 0xf5, //0x0000069d cmpb         $-11, %dl
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x000006a0 ja           LBB0_52
	0x80, 0xe1, 0xdf, //0x000006a6 andb         $-33, %cl
	0x80, 0xc1, 0xb9, //0x000006a9 addb         $-71, %cl
	0x80, 0xf9, 0xfa, //0x000006ac cmpb         $-6, %cl
	0x0f, 0x82, 0x6c, 0xff, 0xff, 0xff, //0x000006af jb           LBB0_101
	//0x000006b5 LBB0_52
	0x48, 0x8d, 0x48, 0x03, //0x000006b5 leaq         $3(%rax), %rcx
	0x48, 0x89, 0x0e, //0x000006b9 movq         %rcx, (%rsi)
	0x41, 0x8a, 0x49, 0x03, //0x000006bc movb         $3(%r9), %cl
	0x8d, 0x51, 0xc6, //0x000006c0 leal         $-58(%rcx), %edx
	0x80, 0xfa, 0xf5, //0x000006c3 cmpb         $-11, %dl
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x000006c6 ja           LBB0_54
	0x80, 0xe1, 0xdf, //0x000006cc andb         $-33, %cl
	0x80, 0xc1, 0xb9, //0x000006cf addb         $-71, %cl
	0x80, 0xf9, 0xfa, //0x000006d2 cmpb         $-6, %cl
	0x0f, 0x82, 0x46, 0xff, 0xff, 0xff, //0x000006d5 jb           LBB0_101
	//0x000006db LBB0_54
	0x48, 0x83, 0xc0, 0x04, //0x000006db addq         $4, %rax
	0x48, 0x89, 0x06, //0x000006df movq         %rax, (%rsi)
	0xe9, 0x3a, 0xff, 0xff, 0xff, //0x000006e2 jmp          LBB0_101
	//0x000006e7 LBB0_39
	0x48, 0x8b, 0x45, 0xc8, //0x000006e7 movq         $-56(%rbp), %rax
	0x48, 0xf7, 0xd0, //0x000006eb notq         %rax
	0x49, 0x01, 0xc1, //0x000006ee addq         %rax, %r9
	0x48, 0x8b, 0x45, 0xd0, //0x000006f1 movq         $-48(%rbp), %rax
	0x4c, 0x89, 0x08, //0x000006f5 movq         %r9, (%rax)
	0x49, 0xc7, 0xc0, 0xfd, 0xff, 0xff, 0xff, //0x000006f8 movq         $-3, %r8
	0xe9, 0x1d, 0xff, 0xff, 0xff, //0x000006ff jmp          LBB0_101
	//0x00000704 LBB0_84
	0x4c, 0x89, 0xc8, //0x00000704 movq         %r9, %rax
	0x48, 0x2b, 0x45, 0xc8, //0x00000707 subq         $-56(%rbp), %rax
	0x48, 0x83, 0xc0, 0x02, //0x0000070b addq         $2, %rax
	0x48, 0x8b, 0x4d, 0xd0, //0x0000070f movq         $-48(%rbp), %rcx
	0x48, 0x89, 0x01, //0x00000713 movq         %rax, (%rcx)
	0x41, 0x8a, 0x49, 0x02, //0x00000716 movb         $2(%r9), %cl
	0x8d, 0x51, 0xc6, //0x0000071a leal         $-58(%rcx), %edx
	0x49, 0xc7, 0xc0, 0xfe, 0xff, 0xff, 0xff, //0x0000071d movq         $-2, %r8
	0x80, 0xfa, 0xf5, //0x00000724 cmpb         $-11, %dl
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x00000727 ja           LBB0_86
	0x80, 0xe1, 0xdf, //0x0000072d andb         $-33, %cl
	0x80, 0xc1, 0xb9, //0x00000730 addb         $-71, %cl
	0x80, 0xf9, 0xfa, //0x00000733 cmpb         $-6, %cl
	0x0f, 0x82, 0xe5, 0xfe, 0xff, 0xff, //0x00000736 jb           LBB0_101
	//0x0000073c LBB0_86
	0x48, 0x8d, 0x48, 0x01, //0x0000073c leaq         $1(%rax), %rcx
	0x48, 0x8b, 0x55, 0xd0, //0x00000740 movq         $-48(%rbp), %rdx
	0x48, 0x89, 0x0a, //0x00000744 movq         %rcx, (%rdx)
	0x41, 0x8a, 0x49, 0x03, //0x00000747 movb         $3(%r9), %cl
	0x8d, 0x51, 0xc6, //0x0000074b leal         $-58(%rcx), %edx
	0x80, 0xfa, 0xf5, //0x0000074e cmpb         $-11, %dl
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x00000751 ja           LBB0_88
	0x80, 0xe1, 0xdf, //0x00000757 andb         $-33, %cl
	0x80, 0xc1, 0xb9, //0x0000075a addb         $-71, %cl
	0x80, 0xf9, 0xfa, //0x0000075d cmpb         $-6, %cl
	0x0f, 0x82, 0xbb, 0xfe, 0xff, 0xff, //0x00000760 jb           LBB0_101
	//0x00000766 LBB0_88
	0x48, 0x8d, 0x48, 0x02, //0x00000766 leaq         $2(%rax), %rcx
	0x48, 0x8b, 0x55, 0xd0, //0x0000076a movq         $-48(%rbp), %rdx
	0x48, 0x89, 0x0a, //0x0000076e movq         %rcx, (%rdx)
	0x41, 0x8a, 0x49, 0x04, //0x00000771 movb         $4(%r9), %cl
	0x8d, 0x51, 0xc6, //0x00000775 leal         $-58(%rcx), %edx
	0x80, 0xfa, 0xf5, //0x00000778 cmpb         $-11, %dl
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x0000077b ja           LBB0_90
	0x80, 0xe1, 0xdf, //0x00000781 andb         $-33, %cl
	0x80, 0xc1, 0xb9, //0x00000784 addb         $-71, %cl
	0x80, 0xf9, 0xfa, //0x00000787 cmpb         $-6, %cl
	0x0f, 0x82, 0x91, 0xfe, 0xff, 0xff, //0x0000078a jb           LBB0_101
	//0x00000790 LBB0_90
	0x48, 0x8d, 0x48, 0x03, //0x00000790 leaq         $3(%rax), %rcx
	0x48, 0x8b, 0x55, 0xd0, //0x00000794 movq         $-48(%rbp), %rdx
	0x48, 0x89, 0x0a, //0x00000798 movq         %rcx, (%rdx)
	0x41, 0x8a, 0x49, 0x05, //0x0000079b movb         $5(%r9), %cl
	0x8d, 0x51, 0xc6, //0x0000079f leal         $-58(%rcx), %edx
	0x80, 0xfa, 0xf5, //0x000007a2 cmpb         $-11, %dl
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x000007a5 ja           LBB0_92
	0x80, 0xe1, 0xdf, //0x000007ab andb         $-33, %cl
	0x80, 0xc1, 0xb9, //0x000007ae addb         $-71, %cl
	0x80, 0xf9, 0xfa, //0x000007b1 cmpb         $-6, %cl
	0x0f, 0x82, 0x67, 0xfe, 0xff, 0xff, //0x000007b4 jb           LBB0_101
	//0x000007ba LBB0_92
	0x48, 0x83, 0xc0, 0x04, //0x000007ba addq         $4, %rax
	0x48, 0x8b, 0x4d, 0xd0, //0x000007be movq         $-48(%rbp), %rcx
	0x48, 0x89, 0x01, //0x000007c2 movq         %rax, (%rcx)
	0xe9, 0x57, 0xfe, 0xff, 0xff, //0x000007c5 jmp          LBB0_101
	//0x000007ca LBB0_28
	0x48, 0x8b, 0x45, 0xc8, //0x000007ca movq         $-56(%rbp), %rax
	0x48, 0xf7, 0xd0, //0x000007ce notq         %rax
	0x49, 0x01, 0xc1, //0x000007d1 addq         %rax, %r9
	//0x000007d4 LBB0_34
	0x48, 0x8b, 0x45, 0xd0, //0x000007d4 movq         $-48(%rbp), %rax
	0x4c, 0x89, 0x08, //0x000007d8 movq         %r9, (%rax)
	0x49, 0xc7, 0xc0, 0xfe, 0xff, 0xff, 0xff, //0x000007db movq         $-2, %r8
	0xe9, 0x3a, 0xfe, 0xff, 0xff, //0x000007e2 jmp          LBB0_101
	//0x000007e7 LBB0_33
	0x4c, 0x2b, 0x4d, 0xc8, //0x000007e7 subq         $-56(%rbp), %r9
	0x49, 0x83, 0xc1, 0x01, //0x000007eb addq         $1, %r9
	0xe9, 0xe0, 0xff, 0xff, 0xff, //0x000007ef jmp          LBB0_34
	//0x000007f4 LBB0_78
	0x4c, 0x2b, 0x4d, 0xc8, //0x000007f4 subq         $-56(%rbp), %r9
	//0x000007f8 LBB0_79
	0x49, 0x83, 0xc1, 0xfc, //0x000007f8 addq         $-4, %r9
	0x48, 0x8b, 0x45, 0xd0, //0x000007fc movq         $-48(%rbp), %rax
	0x4c, 0x89, 0x08, //0x00000800 movq         %r9, (%rax)
	0x49, 0xc7, 0xc0, 0xfc, 0xff, 0xff, 0xff, //0x00000803 movq         $-4, %r8
	0xe9, 0x12, 0xfe, 0xff, 0xff, //0x0000080a jmp          LBB0_101
	//0x0000080f LBB0_75
	0x48, 0x03, 0x75, 0xc8, //0x0000080f addq         $-56(%rbp), %rsi
	0x49, 0x29, 0xf1, //0x00000813 subq         %rsi, %r9
	0xe9, 0xdd, 0xff, 0xff, 0xff, //0x00000816 jmp          LBB0_79
	//0x0000081b LBB0_66
	0xf6, 0x45, 0xb8, 0x02, //0x0000081b testb        $2, $-72(%rbp)
	0x0f, 0x84, 0x1a, 0x00, 0x00, 0x00, //0x0000081f je           LBB0_24
	0x66, 0x41, 0xc7, 0x40, 0xfe, 0xef, 0xbf, //0x00000825 movw         $-16401, $-2(%r8)
	0x41, 0xc6, 0x00, 0xbd, //0x0000082c movb         $-67, (%r8)
	0x49, 0x83, 0xc0, 0x01, //0x00000830 addq         $1, %r8
	0x45, 0x31, 0xd2, //0x00000834 xorl         %r10d, %r10d
	0x4c, 0x89, 0xda, //0x00000837 movq         %r11, %rdx
	0xe9, 0xdc, 0xfd, 0xff, 0xff, //0x0000083a jmp          LBB0_100
	//0x0000083f LBB0_24
	0x48, 0x8b, 0x45, 0xd0, //0x0000083f movq         $-48(%rbp), %rax
	0x48, 0x8b, 0x4d, 0xc0, //0x00000843 movq         $-64(%rbp), %rcx
	0x48, 0x89, 0x08, //0x00000847 movq         %rcx, (%rax)
	0x49, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x0000084a movq         $-1, %r8
	0xe9, 0xcb, 0xfd, 0xff, 0xff, //0x00000851 jmp          LBB0_101
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000856 .p2align 4, 0x00
	//0x00000860 __UnquoteTab
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000860 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .ascii 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000870 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .ascii 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2f, //0x00000880 QUAD $0x0000000000220000; QUAD $0x2f00000000000000  // .ascii 16, '\x00\x00"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00/'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000890 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .ascii 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x000008a0 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .ascii 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5c, 0x00, 0x00, 0x00, //0x000008b0 QUAD $0x0000000000000000; QUAD $0x0000005c00000000  // .ascii 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\\\x00\x00\x00'
	0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, //0x000008c0 QUAD $0x000c000000080000; QUAD $0x000a000000000000  // .ascii 16, '\x00\x00\x08\x00\x00\x00\x0c\x00\x00\x00\x00\x00\x00\x00\n\x00'
	0x00, 0x00, 0x0d, 0x00, 0x09, 0xff, //0x000008d0 LONG $0x000d0000; WORD $0xff09  // .ascii 6, '\x00\x00\r\x00\t\xff'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x000008d6 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .space 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x000008e6 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .space 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x000008f6 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .space 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000906 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .space 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000916 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .space 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000926 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .space 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000936 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .space 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000946 QUAD $0x0000000000000000; QUAD $0x0000000000000000  // .space 16, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000956 QUAD $0x0000000000000000; WORD $0x0000  // .space 10, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
}
