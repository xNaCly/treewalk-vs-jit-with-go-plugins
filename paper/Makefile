all: biber

biber: main
	biber tree-walk-vs-go-jit

main:
	pdflatex -shell-escape -jobname=tree-walk-vs-go-jit main 

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
