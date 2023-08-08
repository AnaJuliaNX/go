#include <stdio.h>
#include <conio.h>
#include <math.h>
#include <locale.h>

main ()
 { setlocale(LC_ALL, "Portuguese");
   int i, j, mat1[2][3], mat2[2][3], mat3[2][3], max;
   
    for (i=0; i<2; i++)
    {
    	for (j=0; j<3; j++)
    	{
		printf("Digite um valor: ");
    	scanf("%i", &mat1[i][j]);
    	
    	 if (mat1[i][j] > max)
   	          {
   	          	max = mat1[i][j];
				 }  
		}
    	
	}
	for (i=0; i<2; i++)
    {
    	for (j=0; j<3; j++)
    	{
		printf("Digite um valor: ");
    	scanf("%i", &mat2[i][j]);
    	
    	if (mat2[i][j] > max)
   	          {
   	          	max = mat2[i][j];
				 }  
		}  	
	}
	
	for (i=0; i<2; i++)
    {
    	for (j=0; j<3; j++)
    	{	
    	 mat3[i][j] = mat1[i][j] + mat2[i][j];
		 
		 printf("\nO resultado será: %i", mat3[i][j]);
		}	
	}
 printf("\nA maior matriz é o: %i", max); 	

}
