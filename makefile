.PHONY: day

dayFolder := day_$(shell date +%d)

day:
	cp -r template/ $(dayFolder)
	@echo making $(dayFolder)
	cd $(dayFolder)