//go:build amd64
// +build amd64

// Code generated by asm2asm, DO NOT EDIT.

package sse

var _text_skip_number = []byte{
	// .p2align 4, 0x00
	// LCPI0_0
	0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, // QUAD $0x2b2b2b2b2b2b2b2b; QUAD $0x2b2b2b2b2b2b2b2b  // .space 16, '++++++++++++++++'
	//0x00000010 LCPI0_1
	0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, //0x00000010 QUAD $0x2d2d2d2d2d2d2d2d; QUAD $0x2d2d2d2d2d2d2d2d  // .space 16, '----------------'
	//0x00000020 LCPI0_2
	0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, //0x00000020 QUAD $0xd0d0d0d0d0d0d0d0; QUAD $0xd0d0d0d0d0d0d0d0  // .space 16, '\xd0\xd0\xd0\xd0\xd0\xd0\xd0\xd0\xd0\xd0\xd0\xd0\xd0\xd0\xd0\xd0'
	//0x00000030 LCPI0_3
	0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, 0x09, //0x00000030 QUAD $0x0909090909090909; QUAD $0x0909090909090909  // .space 16, '\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t'
	//0x00000040 LCPI0_4
	0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, 0xdf, //0x00000040 QUAD $0xdfdfdfdfdfdfdfdf; QUAD $0xdfdfdfdfdfdfdfdf  // .space 16, '\xdf\xdf\xdf\xdf\xdf\xdf\xdf\xdf\xdf\xdf\xdf\xdf\xdf\xdf\xdf\xdf'
	//0x00000050 LCPI0_5
	0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, //0x00000050 QUAD $0x2e2e2e2e2e2e2e2e; QUAD $0x2e2e2e2e2e2e2e2e  // .space 16, '................'
	//0x00000060 LCPI0_6
	0x45, 0x45, 0x45, 0x45, 0x45, 0x45, 0x45, 0x45, 0x45, 0x45, 0x45, 0x45, 0x45, 0x45, 0x45, 0x45, //0x00000060 QUAD $0x4545454545454545; QUAD $0x4545454545454545  // .space 16, 'EEEEEEEEEEEEEEEE'
	//0x00000070 .p2align 4, 0x90
	//0x00000070 _skip_number
	0x55,             //0x00000070 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000071 movq         %rsp, %rbp
	0x41, 0x57, //0x00000074 pushq        %r15
	0x41, 0x56, //0x00000076 pushq        %r14
	0x41, 0x55, //0x00000078 pushq        %r13
	0x41, 0x54, //0x0000007a pushq        %r12
	0x53,                   //0x0000007c pushq        %rbx
	0x48, 0x83, 0xec, 0x18, //0x0000007d subq         $24, %rsp
	0x48, 0x8b, 0x17, //0x00000081 movq         (%rdi), %rdx
	0x4c, 0x8b, 0x6f, 0x08, //0x00000084 movq         $8(%rdi), %r13
	0x4c, 0x8b, 0x0e, //0x00000088 movq         (%rsi), %r9
	0x4d, 0x29, 0xcd, //0x0000008b subq         %r9, %r13
	0x45, 0x31, 0xff, //0x0000008e xorl         %r15d, %r15d
	0x42, 0x80, 0x3c, 0x0a, 0x2d, //0x00000091 cmpb         $45, (%rdx,%r9)
	0x4a, 0x8d, 0x1c, 0x0a, //0x00000096 leaq         (%rdx,%r9), %rbx
	0x41, 0x0f, 0x94, 0xc7, //0x0000009a sete         %r15b
	0x4e, 0x8d, 0x1c, 0x3b, //0x0000009e leaq         (%rbx,%r15), %r11
	0x4d, 0x29, 0xfd, //0x000000a2 subq         %r15, %r13
	0x0f, 0x84, 0xbb, 0x03, 0x00, 0x00, //0x000000a5 je           LBB0_1
	0x41, 0x8a, 0x3b, //0x000000ab movb         (%r11), %dil
	0x8d, 0x4f, 0xc6, //0x000000ae leal         $-58(%rdi), %ecx
	0x48, 0xc7, 0xc0, 0xfe, 0xff, 0xff, 0xff, //0x000000b1 movq         $-2, %rax
	0x80, 0xf9, 0xf6, //0x000000b8 cmpb         $-10, %cl
	0x0f, 0x82, 0x18, 0x03, 0x00, 0x00, //0x000000bb jb           LBB0_59
	0x48, 0x89, 0x55, 0xc8, //0x000000c1 movq         %rdx, $-56(%rbp)
	0x48, 0x89, 0x75, 0xc0, //0x000000c5 movq         %rsi, $-64(%rbp)
	0x40, 0x80, 0xff, 0x30, //0x000000c9 cmpb         $48, %dil
	0x0f, 0x85, 0x34, 0x00, 0x00, 0x00, //0x000000cd jne          LBB0_7
	0xba, 0x01, 0x00, 0x00, 0x00, //0x000000d3 movl         $1, %edx
	0x49, 0x83, 0xfd, 0x01, //0x000000d8 cmpq         $1, %r13
	0x0f, 0x84, 0xcf, 0x02, 0x00, 0x00, //0x000000dc je           LBB0_58
	0x41, 0x8a, 0x43, 0x01, //0x000000e2 movb         $1(%r11), %al
	0x04, 0xd2, //0x000000e6 addb         $-46, %al
	0x3c, 0x37, //0x000000e8 cmpb         $55, %al
	0x0f, 0x87, 0xc1, 0x02, 0x00, 0x00, //0x000000ea ja           LBB0_58
	0x0f, 0xb6, 0xc0, //0x000000f0 movzbl       %al, %eax
	0x48, 0xb9, 0x01, 0x00, 0x80, 0x00, 0x00, 0x00, 0x80, 0x00, //0x000000f3 movabsq      $36028797027352577, %rcx
	0x48, 0x0f, 0xa3, 0xc1, //0x000000fd btq          %rax, %rcx
	0x0f, 0x83, 0xaa, 0x02, 0x00, 0x00, //0x00000101 jae          LBB0_58
	//0x00000107 LBB0_7
	0x48, 0x89, 0x5d, 0xd0, //0x00000107 movq         %rbx, $-48(%rbp)
	0x49, 0x83, 0xfd, 0x10, //0x0000010b cmpq         $16, %r13
	0x0f, 0x82, 0x5d, 0x03, 0x00, 0x00, //0x0000010f jb           LBB0_8
	0x49, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x00000115 movq         $-1, %r8
	0x31, 0xd2, //0x0000011c xorl         %edx, %edx
	0xf3, 0x44, 0x0f, 0x6f, 0x05, 0xd9, 0xfe, 0xff, 0xff, //0x0000011e movdqu       $-295(%rip), %xmm8  /* LCPI0_0+0(%rip) */
	0xf3, 0x44, 0x0f, 0x6f, 0x0d, 0xe0, 0xfe, 0xff, 0xff, //0x00000127 movdqu       $-288(%rip), %xmm9  /* LCPI0_1+0(%rip) */
	0xf3, 0x44, 0x0f, 0x6f, 0x15, 0xe7, 0xfe, 0xff, 0xff, //0x00000130 movdqu       $-281(%rip), %xmm10  /* LCPI0_2+0(%rip) */
	0xf3, 0x0f, 0x6f, 0x1d, 0xef, 0xfe, 0xff, 0xff, //0x00000139 movdqu       $-273(%rip), %xmm3  /* LCPI0_3+0(%rip) */
	0xf3, 0x0f, 0x6f, 0x25, 0xf7, 0xfe, 0xff, 0xff, //0x00000141 movdqu       $-265(%rip), %xmm4  /* LCPI0_4+0(%rip) */
	0xf3, 0x0f, 0x6f, 0x2d, 0xff, 0xfe, 0xff, 0xff, //0x00000149 movdqu       $-257(%rip), %xmm5  /* LCPI0_5+0(%rip) */
	0xf3, 0x0f, 0x6f, 0x35, 0x07, 0xff, 0xff, 0xff, //0x00000151 movdqu       $-249(%rip), %xmm6  /* LCPI0_6+0(%rip) */
	0x49, 0xc7, 0xc4, 0xff, 0xff, 0xff, 0xff, //0x00000159 movq         $-1, %r12
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x00000160 movq         $-1, %r14
	0x4c, 0x89, 0xef, //0x00000167 movq         %r13, %rdi
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x0000016a .p2align 4, 0x90
	//0x00000170 LBB0_10
	0xf3, 0x41, 0x0f, 0x6f, 0x3c, 0x13, //0x00000170 movdqu       (%r11,%rdx), %xmm7
	0x66, 0x0f, 0x6f, 0xc7, //0x00000176 movdqa       %xmm7, %xmm0
	0x66, 0x41, 0x0f, 0x74, 0xc0, //0x0000017a pcmpeqb      %xmm8, %xmm0
	0x66, 0x0f, 0x6f, 0xcf, //0x0000017f movdqa       %xmm7, %xmm1
	0x66, 0x41, 0x0f, 0x74, 0xc9, //0x00000183 pcmpeqb      %xmm9, %xmm1
	0x66, 0x0f, 0xeb, 0xc8, //0x00000188 por          %xmm0, %xmm1
	0x66, 0x0f, 0x6f, 0xc7, //0x0000018c movdqa       %xmm7, %xmm0
	0x66, 0x41, 0x0f, 0xfc, 0xc2, //0x00000190 paddb        %xmm10, %xmm0
	0x66, 0x0f, 0x6f, 0xd0, //0x00000195 movdqa       %xmm0, %xmm2
	0x66, 0x0f, 0xda, 0xd3, //0x00000199 pminub       %xmm3, %xmm2
	0x66, 0x0f, 0x74, 0xd0, //0x0000019d pcmpeqb      %xmm0, %xmm2
	0x66, 0x0f, 0x6f, 0xc7, //0x000001a1 movdqa       %xmm7, %xmm0
	0x66, 0x0f, 0xdb, 0xc4, //0x000001a5 pand         %xmm4, %xmm0
	0x66, 0x0f, 0x74, 0xc6, //0x000001a9 pcmpeqb      %xmm6, %xmm0
	0x66, 0x0f, 0x74, 0xfd, //0x000001ad pcmpeqb      %xmm5, %xmm7
	0x66, 0x0f, 0xd7, 0xf0, //0x000001b1 pmovmskb     %xmm0, %esi
	0x66, 0x0f, 0xeb, 0xc7, //0x000001b5 por          %xmm7, %xmm0
	0x66, 0x0f, 0xeb, 0xc1, //0x000001b9 por          %xmm1, %xmm0
	0x66, 0x0f, 0xeb, 0xc2, //0x000001bd por          %xmm2, %xmm0
	0x66, 0x0f, 0xd7, 0xdf, //0x000001c1 pmovmskb     %xmm7, %ebx
	0x66, 0x44, 0x0f, 0xd7, 0xd1, //0x000001c5 pmovmskb     %xmm1, %r10d
	0x66, 0x0f, 0xd7, 0xc0, //0x000001ca pmovmskb     %xmm0, %eax
	0xf7, 0xd0, //0x000001ce notl         %eax
	0x0f, 0xbc, 0xc8, //0x000001d0 bsfl         %eax, %ecx
	0x83, 0xf9, 0x10, //0x000001d3 cmpl         $16, %ecx
	0x0f, 0x84, 0x13, 0x00, 0x00, 0x00, //0x000001d6 je           LBB0_12
	0xb8, 0xff, 0xff, 0xff, 0xff, //0x000001dc movl         $-1, %eax
	0xd3, 0xe0, //0x000001e1 shll         %cl, %eax
	0xf7, 0xd0, //0x000001e3 notl         %eax
	0x21, 0xc3, //0x000001e5 andl         %eax, %ebx
	0x21, 0xc6, //0x000001e7 andl         %eax, %esi
	0x44, 0x21, 0xd0, //0x000001e9 andl         %r10d, %eax
	0x41, 0x89, 0xc2, //0x000001ec movl         %eax, %r10d
	//0x000001ef LBB0_12
	0x8d, 0x43, 0xff, //0x000001ef leal         $-1(%rbx), %eax
	0x21, 0xd8, //0x000001f2 andl         %ebx, %eax
	0x0f, 0x85, 0x34, 0x02, 0x00, 0x00, //0x000001f4 jne          LBB0_13
	0x8d, 0x46, 0xff, //0x000001fa leal         $-1(%rsi), %eax
	0x21, 0xf0, //0x000001fd andl         %esi, %eax
	0x0f, 0x85, 0x29, 0x02, 0x00, 0x00, //0x000001ff jne          LBB0_13
	0x41, 0x8d, 0x42, 0xff, //0x00000205 leal         $-1(%r10), %eax
	0x44, 0x21, 0xd0, //0x00000209 andl         %r10d, %eax
	0x0f, 0x85, 0x1c, 0x02, 0x00, 0x00, //0x0000020c jne          LBB0_13
	0x85, 0xdb, //0x00000212 testl        %ebx, %ebx
	0x0f, 0x84, 0x13, 0x00, 0x00, 0x00, //0x00000214 je           LBB0_20
	0x0f, 0xbc, 0xdb, //0x0000021a bsfl         %ebx, %ebx
	0x49, 0x83, 0xfe, 0xff, //0x0000021d cmpq         $-1, %r14
	0x0f, 0x85, 0x2a, 0x02, 0x00, 0x00, //0x00000221 jne          LBB0_60
	0x48, 0x01, 0xd3, //0x00000227 addq         %rdx, %rbx
	0x49, 0x89, 0xde, //0x0000022a movq         %rbx, %r14
	//0x0000022d LBB0_20
	0x85, 0xf6, //0x0000022d testl        %esi, %esi
	0x0f, 0x84, 0x13, 0x00, 0x00, 0x00, //0x0000022f je           LBB0_23
	0x0f, 0xbc, 0xf6, //0x00000235 bsfl         %esi, %esi
	0x49, 0x83, 0xfc, 0xff, //0x00000238 cmpq         $-1, %r12
	0x0f, 0x85, 0x16, 0x02, 0x00, 0x00, //0x0000023c jne          LBB0_61
	0x48, 0x01, 0xd6, //0x00000242 addq         %rdx, %rsi
	0x49, 0x89, 0xf4, //0x00000245 movq         %rsi, %r12
	//0x00000248 LBB0_23
	0x45, 0x85, 0xd2, //0x00000248 testl        %r10d, %r10d
	0x0f, 0x84, 0x14, 0x00, 0x00, 0x00, //0x0000024b je           LBB0_26
	0x41, 0x0f, 0xbc, 0xc2, //0x00000251 bsfl         %r10d, %eax
	0x49, 0x83, 0xf8, 0xff, //0x00000255 cmpq         $-1, %r8
	0x0f, 0x85, 0x00, 0x02, 0x00, 0x00, //0x00000259 jne          LBB0_62
	0x48, 0x01, 0xd0, //0x0000025f addq         %rdx, %rax
	0x49, 0x89, 0xc0, //0x00000262 movq         %rax, %r8
	//0x00000265 LBB0_26
	0x83, 0xf9, 0x10, //0x00000265 cmpl         $16, %ecx
	0x0f, 0x85, 0xab, 0x00, 0x00, 0x00, //0x00000268 jne          LBB0_63
	0x48, 0x83, 0xc7, 0xf0, //0x0000026e addq         $-16, %rdi
	0x48, 0x83, 0xc2, 0x10, //0x00000272 addq         $16, %rdx
	0x48, 0x83, 0xff, 0x0f, //0x00000276 cmpq         $15, %rdi
	0x0f, 0x87, 0xf0, 0xfe, 0xff, 0xff, //0x0000027a ja           LBB0_10
	0x49, 0x8d, 0x0c, 0x13, //0x00000280 leaq         (%r11,%rdx), %rcx
	0x49, 0x89, 0xca, //0x00000284 movq         %rcx, %r10
	0x49, 0x39, 0xd5, //0x00000287 cmpq         %rdx, %r13
	0x0f, 0x84, 0x92, 0x00, 0x00, 0x00, //0x0000028a je           LBB0_42
	//0x00000290 LBB0_29
	0x4c, 0x8d, 0x14, 0x39, //0x00000290 leaq         (%rcx,%rdi), %r10
	0x48, 0x89, 0xc8, //0x00000294 movq         %rcx, %rax
	0x4c, 0x29, 0xd8, //0x00000297 subq         %r11, %rax
	0x31, 0xd2, //0x0000029a xorl         %edx, %edx
	0x4c, 0x8d, 0x2d, 0xf1, 0x01, 0x00, 0x00, //0x0000029c leaq         $497(%rip), %r13  /* LJTI0_0+0(%rip) */
	0xe9, 0x25, 0x00, 0x00, 0x00, //0x000002a3 jmp          LBB0_30
	//0x000002a8 LBB0_32
	0x83, 0xfb, 0x65, //0x000002a8 cmpl         $101, %ebx
	0x0f, 0x85, 0x86, 0x00, 0x00, 0x00, //0x000002ab jne          LBB0_41
	//0x000002b1 LBB0_33
	0x49, 0x83, 0xfc, 0xff, //0x000002b1 cmpq         $-1, %r12
	0x0f, 0x85, 0x81, 0x01, 0x00, 0x00, //0x000002b5 jne          LBB0_64
	0x4c, 0x8d, 0x24, 0x10, //0x000002bb leaq         (%rax,%rdx), %r12
	0x90, //0x000002bf .p2align 4, 0x90
	//0x000002c0 LBB0_40
	0x48, 0x83, 0xc2, 0x01, //0x000002c0 addq         $1, %rdx
	0x48, 0x39, 0xd7, //0x000002c4 cmpq         %rdx, %rdi
	0x0f, 0x84, 0x55, 0x00, 0x00, 0x00, //0x000002c7 je           LBB0_42
	//0x000002cd LBB0_30
	0x0f, 0xbe, 0x1c, 0x11, //0x000002cd movsbl       (%rcx,%rdx), %ebx
	0x8d, 0x73, 0xd0, //0x000002d1 leal         $-48(%rbx), %esi
	0x83, 0xfe, 0x0a, //0x000002d4 cmpl         $10, %esi
	0x0f, 0x82, 0xe3, 0xff, 0xff, 0xff, //0x000002d7 jb           LBB0_40
	0x8d, 0x73, 0xd5, //0x000002dd leal         $-43(%rbx), %esi
	0x83, 0xfe, 0x1a, //0x000002e0 cmpl         $26, %esi
	0x0f, 0x87, 0xbf, 0xff, 0xff, 0xff, //0x000002e3 ja           LBB0_32
	0x49, 0x63, 0x74, 0xb5, 0x00, //0x000002e9 movslq       (%r13,%rsi,4), %rsi
	0x4c, 0x01, 0xee, //0x000002ee addq         %r13, %rsi
	0xff, 0xe6, //0x000002f1 jmpq         *%rsi
	//0x000002f3 LBB0_38
	0x49, 0x83, 0xf8, 0xff, //0x000002f3 cmpq         $-1, %r8
	0x0f, 0x85, 0x3f, 0x01, 0x00, 0x00, //0x000002f7 jne          LBB0_64
	0x4c, 0x8d, 0x04, 0x10, //0x000002fd leaq         (%rax,%rdx), %r8
	0xe9, 0xba, 0xff, 0xff, 0xff, //0x00000301 jmp          LBB0_40
	//0x00000306 LBB0_36
	0x49, 0x83, 0xfe, 0xff, //0x00000306 cmpq         $-1, %r14
	0x0f, 0x85, 0x2c, 0x01, 0x00, 0x00, //0x0000030a jne          LBB0_64
	0x4c, 0x8d, 0x34, 0x10, //0x00000310 leaq         (%rax,%rdx), %r14
	0xe9, 0xa7, 0xff, 0xff, 0xff, //0x00000314 jmp          LBB0_40
	//0x00000319 LBB0_63
	0x41, 0x89, 0xca, //0x00000319 movl         %ecx, %r10d
	0x4d, 0x01, 0xda, //0x0000031c addq         %r11, %r10
	0x49, 0x01, 0xd2, //0x0000031f addq         %rdx, %r10
	//0x00000322 LBB0_42
	0x48, 0xc7, 0xc2, 0xff, 0xff, 0xff, 0xff, //0x00000322 movq         $-1, %rdx
	0x4d, 0x85, 0xf6, //0x00000329 testq        %r14, %r14
	0x0f, 0x85, 0x1b, 0x00, 0x00, 0x00, //0x0000032c jne          LBB0_43
	0xe9, 0x8d, 0x00, 0x00, 0x00, //0x00000332 jmp          LBB0_57
	//0x00000337 LBB0_41
	0x48, 0x01, 0xd1, //0x00000337 addq         %rdx, %rcx
	0x49, 0x89, 0xca, //0x0000033a movq         %rcx, %r10
	0x48, 0xc7, 0xc2, 0xff, 0xff, 0xff, 0xff, //0x0000033d movq         $-1, %rdx
	0x4d, 0x85, 0xf6, //0x00000344 testq        %r14, %r14
	0x0f, 0x84, 0x77, 0x00, 0x00, 0x00, //0x00000347 je           LBB0_57
	//0x0000034d LBB0_43
	0x4d, 0x85, 0xc0, //0x0000034d testq        %r8, %r8
	0x0f, 0x84, 0x6e, 0x00, 0x00, 0x00, //0x00000350 je           LBB0_57
	0x4d, 0x85, 0xe4, //0x00000356 testq        %r12, %r12
	0x0f, 0x84, 0x65, 0x00, 0x00, 0x00, //0x00000359 je           LBB0_57
	0x4d, 0x29, 0xda, //0x0000035f subq         %r11, %r10
	0x49, 0x8d, 0x42, 0xff, //0x00000362 leaq         $-1(%r10), %rax
	0x49, 0x39, 0xc6, //0x00000366 cmpq         %rax, %r14
	0x0f, 0x84, 0x33, 0x00, 0x00, 0x00, //0x00000369 je           LBB0_48
	0x49, 0x39, 0xc0, //0x0000036f cmpq         %rax, %r8
	0x0f, 0x84, 0x2a, 0x00, 0x00, 0x00, //0x00000372 je           LBB0_48
	0x49, 0x39, 0xc4, //0x00000378 cmpq         %rax, %r12
	0x0f, 0x84, 0x21, 0x00, 0x00, 0x00, //0x0000037b je           LBB0_48
	0x4d, 0x85, 0xc0, //0x00000381 testq        %r8, %r8
	0x0f, 0x8e, 0x64, 0x00, 0x00, 0x00, //0x00000384 jle          LBB0_52
	0x49, 0x8d, 0x40, 0xff, //0x0000038a leaq         $-1(%r8), %rax
	0x49, 0x39, 0xc4, //0x0000038e cmpq         %rax, %r12
	0x0f, 0x84, 0x57, 0x00, 0x00, 0x00, //0x00000391 je           LBB0_52
	0x49, 0xf7, 0xd0, //0x00000397 notq         %r8
	0x4c, 0x89, 0xc2, //0x0000039a movq         %r8, %rdx
	0xe9, 0x06, 0x00, 0x00, 0x00, //0x0000039d jmp          LBB0_56
	//0x000003a2 LBB0_48
	0x49, 0xf7, 0xda, //0x000003a2 negq         %r10
	0x4c, 0x89, 0xd2, //0x000003a5 movq         %r10, %rdx
	//0x000003a8 LBB0_56
	0x48, 0x85, 0xd2, //0x000003a8 testq        %rdx, %rdx
	0x0f, 0x88, 0x13, 0x00, 0x00, 0x00, //0x000003ab js           LBB0_57
	//0x000003b1 LBB0_58
	0x49, 0x01, 0xd3, //0x000003b1 addq         %rdx, %r11
	0x4c, 0x89, 0xc8, //0x000003b4 movq         %r9, %rax
	0x48, 0x8b, 0x75, 0xc0, //0x000003b7 movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x55, 0xc8, //0x000003bb movq         $-56(%rbp), %rdx
	0xe9, 0x15, 0x00, 0x00, 0x00, //0x000003bf jmp          LBB0_59
	//0x000003c4 LBB0_57
	0x48, 0xf7, 0xd2, //0x000003c4 notq         %rdx
	0x49, 0x01, 0xd3, //0x000003c7 addq         %rdx, %r11
	0x48, 0x8b, 0x75, 0xc0, //0x000003ca movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x55, 0xc8, //0x000003ce movq         $-56(%rbp), %rdx
	0x48, 0xc7, 0xc0, 0xfe, 0xff, 0xff, 0xff, //0x000003d2 movq         $-2, %rax
	//0x000003d9 LBB0_59
	0x49, 0x29, 0xd3, //0x000003d9 subq         %rdx, %r11
	0x4c, 0x89, 0x1e, //0x000003dc movq         %r11, (%rsi)
	0x48, 0x83, 0xc4, 0x18, //0x000003df addq         $24, %rsp
	0x5b,       //0x000003e3 popq         %rbx
	0x41, 0x5c, //0x000003e4 popq         %r12
	0x41, 0x5d, //0x000003e6 popq         %r13
	0x41, 0x5e, //0x000003e8 popq         %r14
	0x41, 0x5f, //0x000003ea popq         %r15
	0x5d, //0x000003ec popq         %rbp
	0xc3, //0x000003ed retq
	//0x000003ee LBB0_52
	0x4c, 0x89, 0xf0, //0x000003ee movq         %r14, %rax
	0x4c, 0x09, 0xe0, //0x000003f1 orq          %r12, %rax
	0x0f, 0x99, 0xc0, //0x000003f4 setns        %al
	0x0f, 0x88, 0x14, 0x00, 0x00, 0x00, //0x000003f7 js           LBB0_55
	0x4d, 0x39, 0xe6, //0x000003fd cmpq         %r12, %r14
	0x0f, 0x8c, 0x0b, 0x00, 0x00, 0x00, //0x00000400 jl           LBB0_55
	0x49, 0xf7, 0xd6, //0x00000406 notq         %r14
	0x4c, 0x89, 0xf2, //0x00000409 movq         %r14, %rdx
	0xe9, 0x97, 0xff, 0xff, 0xff, //0x0000040c jmp          LBB0_56
	//0x00000411 LBB0_55
	0x49, 0x8d, 0x4c, 0x24, 0xff, //0x00000411 leaq         $-1(%r12), %rcx
	0x49, 0x39, 0xce, //0x00000416 cmpq         %rcx, %r14
	0x49, 0xf7, 0xd4, //0x00000419 notq         %r12
	0x4d, 0x0f, 0x45, 0xe2, //0x0000041c cmovneq      %r10, %r12
	0x84, 0xc0, //0x00000420 testb        %al, %al
	0x4d, 0x0f, 0x44, 0xe2, //0x00000422 cmoveq       %r10, %r12
	0x4c, 0x89, 0xe2, //0x00000426 movq         %r12, %rdx
	0xe9, 0x7a, 0xff, 0xff, 0xff, //0x00000429 jmp          LBB0_56
	//0x0000042e LBB0_13
	0x0f, 0xbc, 0xc0, //0x0000042e bsfl         %eax, %eax
	//0x00000431 LBB0_14
	0x48, 0xf7, 0xd2, //0x00000431 notq         %rdx
	0x48, 0x29, 0xc2, //0x00000434 subq         %rax, %rdx
	0xe9, 0x6c, 0xff, 0xff, 0xff, //0x00000437 jmp          LBB0_56
	//0x0000043c LBB0_64
	0x48, 0x8b, 0x45, 0xd0, //0x0000043c movq         $-48(%rbp), %rax
	0x4c, 0x01, 0xf8, //0x00000440 addq         %r15, %rax
	0x48, 0x29, 0xc8, //0x00000443 subq         %rcx, %rax
	0x48, 0xf7, 0xd2, //0x00000446 notq         %rdx
	0x48, 0x01, 0xc2, //0x00000449 addq         %rax, %rdx
	0xe9, 0x57, 0xff, 0xff, 0xff, //0x0000044c jmp          LBB0_56
	//0x00000451 LBB0_60
	0x89, 0xd8, //0x00000451 movl         %ebx, %eax
	0xe9, 0xd9, 0xff, 0xff, 0xff, //0x00000453 jmp          LBB0_14
	//0x00000458 LBB0_61
	0x89, 0xf0, //0x00000458 movl         %esi, %eax
	0xe9, 0xd2, 0xff, 0xff, 0xff, //0x0000045a jmp          LBB0_14
	//0x0000045f LBB0_62
	0x89, 0xc0, //0x0000045f movl         %eax, %eax
	0xe9, 0xcb, 0xff, 0xff, 0xff, //0x00000461 jmp          LBB0_14
	//0x00000466 LBB0_1
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x00000466 movq         $-1, %rax
	0xe9, 0x67, 0xff, 0xff, 0xff, //0x0000046d jmp          LBB0_59
	//0x00000472 LBB0_8
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x00000472 movq         $-1, %r14
	0x4c, 0x89, 0xd9, //0x00000479 movq         %r11, %rcx
	0x4c, 0x89, 0xef, //0x0000047c movq         %r13, %rdi
	0x49, 0xc7, 0xc4, 0xff, 0xff, 0xff, 0xff, //0x0000047f movq         $-1, %r12
	0x49, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x00000486 movq         $-1, %r8
	0xe9, 0xfe, 0xfd, 0xff, 0xff, //0x0000048d jmp          LBB0_29
	0x90, 0x90, //0x00000492 .p2align 2, 0x90
	// // .set L0_0_set_38, LBB0_38-LJTI0_0
	// // .set L0_0_set_41, LBB0_41-LJTI0_0
	// // .set L0_0_set_36, LBB0_36-LJTI0_0
	// // .set L0_0_set_33, LBB0_33-LJTI0_0
	//0x00000494 LJTI0_0
	0x5f, 0xfe, 0xff, 0xff, //0x00000494 .long L0_0_set_38
	0xa3, 0xfe, 0xff, 0xff, //0x00000498 .long L0_0_set_41
	0x5f, 0xfe, 0xff, 0xff, //0x0000049c .long L0_0_set_38
	0x72, 0xfe, 0xff, 0xff, //0x000004a0 .long L0_0_set_36
	0xa3, 0xfe, 0xff, 0xff, //0x000004a4 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004a8 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004ac .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004b0 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004b4 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004b8 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004bc .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004c0 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004c4 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004c8 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004cc .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004d0 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004d4 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004d8 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004dc .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004e0 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004e4 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004e8 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004ec .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004f0 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004f4 .long L0_0_set_41
	0xa3, 0xfe, 0xff, 0xff, //0x000004f8 .long L0_0_set_41
	0x1d, 0xfe, 0xff, 0xff, //0x000004fc .long L0_0_set_33
	//0x00000500 .p2align 2, 0x00
	//0x00000500 _MASK_USE_NUMBER
	0x02, 0x00, 0x00, 0x00, //0x00000500 .long 2
}