section .data
    heart db "❤️"

section .text
    global _start
_start:
    mov eax, 4
    mov ebx, 1
    mov ecx, heart
    mov edx, 3
loop:
    int 0x80
    jmp loop