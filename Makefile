# Makefile

# Clean job to delete all *resized.png files
clean:
	@echo "Deleting *resized.png files..."
	@find . -name '*resized.png' -delete
	@echo "Clean job completed."

