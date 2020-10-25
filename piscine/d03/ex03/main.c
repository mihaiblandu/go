/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.c                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mihai <mihai@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/10/22 22:51:33 by mihai             #+#    #+#             */
/*   Updated: 2020/10/22 22:59:32 by mihai            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <stdio.h>

void ft_ultimate_div_mod(int *a, int *b);

int main(){
    int i = 0;
    int j = 0;
	int *a = &i;
	int *b = &j;

	ft_div_mod(3, 25, a, b);
	
	printf("\nValue of i is : %d",*a);
	printf("\nValue of j is : %d",*b);
}