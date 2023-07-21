resource_methods.html: resource_methods.org diag1.png
	pandoc --standalone --slide-level=2 --to revealjs --css style.css \
		-o $@ $<

diag1.png: diag1.mmd
	mmdc -i diag1.mmd -o diag1.png
