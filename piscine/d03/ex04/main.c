/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.c                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mihai <mihai@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/10/22 23:06:54 by mihai             #+#    #+#             */
/*   Updated: 2020/10/22 23:06:55 by mihai            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <stdio.h>

void ft_ultimate_div_mod(int *a, int *b);

int main(){
	int i = -10;
	int j = 3;
	int *a = &i;
	int *b = &j;

	ft_ultimate_div_mod(a, b);
	
	printf("\nValue of i is : %d",*a);
	printf("\nValue of j is : %d",*b);
}