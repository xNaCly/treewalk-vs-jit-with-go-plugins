all: biber

biber: main
	biber main

main:
	pdflatex -shell-escape main

clean:
	rm -fr \
		*.toc \
		**/**.aux \
		*.aux \
		*.nav \
		*.out \
		*.snm \
		*.toc \
		*.log \
		*.bbl \
		*.blg \
		*.bcf \
		*.run.xml
