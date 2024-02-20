addi r1 r0 0 # start addr
addi r2 r0 63 # add 63 to r3
addi r3 r0 3 # jump addr
sw r1 r1 0 # load r1 into the value of r1
beq r1 r2 2 # if r1 and r2 are equal jump past jalr instuction
addi r1 r1 1 # add 1 to r1
jalr r0 r3 # jump to value of r3