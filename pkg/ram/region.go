package ram

func BIOS(addr uint32) (bool, uint32) {
	return addr < 0x0200_0000, addr
}
func EWRAM(addr uint32) (bool, uint32) {
	return 0x0200_0000 <= addr && addr < 0x0300_0000, addr - 0x0200_0000
}
func IWRAM(addr uint32) (bool, uint32) {
	return 0x0300_0000 <= addr && addr < 0x0400_0000, addr - 0x0300_0000
}
func IO(addr uint32) (bool, uint32) {
	return 0x0400_0000 <= addr && addr < 0x0500_0000, addr - 0x0400_0000
}
func Palette(addr uint32) (bool, uint32) {
	return 0x0500_0000 <= addr && addr < 0x0600_0000, addr - 0x0500_0000
}
func VRAM(addr uint32) (bool, uint32) {
	return 0x0600_0000 <= addr && addr < 0x0700_0000, addr - 0x0600_0000
}
func OAM(addr uint32) (bool, uint32) {
	return 0x0700_0000 <= addr && addr < 0x0800_0000, addr - 0x0700_0000
}
func GamePak0(addr uint32) (bool, uint32) {
	return 0x0800_0000 <= addr && addr < 0x0a00_0000, addr - 0x0800_0000
}
func GamePak1(addr uint32) (bool, uint32) {
	return 0x0a00_0000 <= addr && addr < 0x0c00_0000, addr - 0x0a00_0000
}
func GamePak2(addr uint32) (bool, uint32) {
	return 0x0c00_0000 <= addr && addr < 0x0e00_0000, addr - 0x0c00_0000
}
func SRAM(addr uint32) (bool, uint32) {
	return 0x0e0_0000 <= addr && addr < 0x0e01_0000, addr - 0x0e0_0000
}
